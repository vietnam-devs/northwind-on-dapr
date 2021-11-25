using System.Net;
using System.Text.Json;
using Dapr.Client;
using Grpc.Net.ClientFactory;
using Northwind.Protobuf.Product;

var builder = WebApplication.CreateBuilder(args);

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

app.UseHttpsRedirection();

app.UseCloudEvents();

app.MapSubscribeHandler();

app.MapFallback(() => Results.Redirect("/swagger"));

app.MapGet("/ping", () => Results.Ok("Okay"))
    .WithName("GetWeatherForecast");

app.MapGet("/api/products", async (GrpcClientFactory grpcClientFactory) => {
    var productClient = grpcClientFactory.CreateClient<ProductApi.ProductApiClient>("product-client");
    var result = await productClient.GetProductsAsync(new GetProductsRequest());
    return Results.Ok(result);
});

// app.MapGet("/api/products", async (DaprClient daprClient) => {
    // call Dapr get products
// });

app.MapPost("/v1/order", async (DaprClient client) => {
    // direct call Dapr get products - product service
    // pubsub Kafka - shipping service
    
    //var client1 = DaprClient.CreateInvokeHttpClient(appId: "shipping");
    //var result = await client1.GetStringAsync("/");
    await client.PublishEventAsync("pubsub", "order", new OrderCreated(Guid.NewGuid()));
});

app.Run();

public record struct OrderCreated(Guid Id);