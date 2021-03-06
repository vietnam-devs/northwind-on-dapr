using System.Text.Json;
using Dapr.Client;
using Grpc.Core;
using Grpc.Net.Client;
using Northwind.Protobuf.Product;

var builder = WebApplication.CreateBuilder(args);

builder.Logging.AddJsonConsole();

// Add services to the container.
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

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

app.MapGet("/api/products", async () =>
{
    var port = Environment.GetEnvironmentVariable("DAPR_GRPC_PORT") ?? "50003";
    var serverAddress = $"http://localhost:{port}";
    var channel = GrpcChannel.ForAddress(serverAddress);
    var client = new ProductApi.ProductApiClient(channel);

    var metaData = new Metadata { { "dapr-app-id", "product-catalog" } };
    var result = await client.GetProductsAsync(new GetProductsRequest(), metaData);
    return Results.Ok(result);
});

app.MapGet("/api/shipping", async (DaprClient client) =>
{
    var result = await client.InvokeMethodAsync<object>(httpMethod: HttpMethod.Get, "shipping", "/");
    return Results.Ok(result);
});

app.MapPost("/v1/order", async (DaprClient client) =>
{
    // direct call Dapr get products - product service
    // pubsub Kafka - shipping service

    await client.PublishEventAsync("pubsub", "order", new OrderCreated(Guid.NewGuid()));
});

// app.MapPost("/api/v1/subscribers/order-created", (OrderCreated message) =>
//     {
//         app.Logger.LogInformation("Received message");
//         return Results.Ok(true);
//     })
//     .WithTopic("pubsub", "order");

app.Run();

public record OrderCreated(Guid Id);
