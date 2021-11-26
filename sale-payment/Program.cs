using System.Text.Json;
using Dapr.Client;
using Grpc.Net.ClientFactory;
using Northwind.Protobuf.Product;

var builder = WebApplication.CreateBuilder(args);

builder.Logging.AddJsonConsole();

// Add services to the container.
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

builder.Services.AddGrpcClient<ProductApi.ProductApiClient>("product-client", o =>
    {
        o.Address = new Uri(builder.Configuration.GetValue<string>("ProductGrpcUrl"));
    })
    .EnableCallContextPropagation(o => o.SuppressContextNotFoundErrors = true);

builder.Services.AddDaprClient(builder =>
    builder.UseJsonSerializationOptions(
        new JsonSerializerOptions()
        {
            PropertyNamingPolicy = JsonNamingPolicy.CamelCase,
            PropertyNameCaseInsensitive = true,
        }));

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseCloudEvents();

app.MapSubscribeHandler();

app.MapFallback(() => Results.Redirect("/swagger"));

app.MapGet("/ping", () => Results.Ok("Okay"))
    .WithName("GetWeatherForecast");

app.MapGet("/api/products", async (GrpcClientFactory grpcClientFactory) =>
{
    var productClient = grpcClientFactory.CreateClient<ProductApi.ProductApiClient>("product-client");
    var result = await productClient.GetProductsAsync(new GetProductsRequest());
    return Results.Ok(result);
});

// app.MapGet("/api/products", async () => {
//     var client = DaprClient.CreateInvokeHttpClient(appId: "product-catalog");
//     var result = await client.GetStringAsync("/v1/products");
//     return Results.Ok(result);
// });

app.MapPost("/v1/order", async (DaprClient client) =>
{
    // direct call Dapr get products - product service
    // pubsub Kafka - shipping service

    await client.PublishEventAsync("pubsub", "order", new OrderCreated(Guid.NewGuid()));
});

app.MapPost("/api/v1/subscribers/order-created", (OrderCreated message) =>
    {
        app.Logger.LogInformation("Received message");
        return Results.Ok(true);
    })
    .WithTopic("pubsub", "order");

app.Run();

public record OrderCreated(Guid Id);
