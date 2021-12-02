# northwind-on-dapr
This is a demonstration for how to use Dapr with polyglot programming approach

# Setup env

- Add `.env` file on the root of this repository. That means on https://github.com/vietnam-devs/northwind-on-dapr

```
PRODUCT_HOST=localhost # if on Docker then it should be 0.0.0.0

POSTGRES_USER=northwind
POSTGRES_PASSWORD=<your password>
POSTGRES_DB=northwind_db

ProductGrpcUrl=http://localhost:50002
```
## Local run - using Tye only

```bash
> tye run
```
## Local run and debug manually

- Run external services using Tye
```bash
> tye run --tags inf
```
Once external services such as postgres, rabbitMQ,... are started, continue to next commmands to launch microservices
- Run the `go-app` (product-catalog) service
```bash
> cd ./product-catalog
> dapr run --app-id product-catalog --app-port 50002 --components-path ..\components\ --config ..\components\config.yaml -- go run .
```
- Run the `dotnet-core-app` (sale-payment) service
```bash
> cd ./sale-payment
> dapr run --app-id sale-payment --app-port 5003 --dapr-grpc-port 50003 --components-path ..\components\ --config ..\components\config.yaml -- dotnet watch run
```
- Run the `java-app` (shipping) service
```bash
> cd ./sale-payment
> dapr run --app-id shipping --app-port 5004 --components-path ..\components\ --config ..\components\config.yaml -- mvn spring-boot:run
```

# Invoke API using Dapr CLI
## `sale-payment` service
- Invoke `/ping`
```bash
dapr invoke --app-id sale-payment -m /ping -v Get
```
- Get all products of `product-catalog` service via `sale-payment` service using grpc proxy feature
```bash
dapr invoke --app-id sale-payment -m /api/products -v Get
```
## `shipping` service
- Invoke `/`
```bash
dapr invoke --app-id shipping -m / -v Get
```

# Observability
## Distributed tracing
It is enabled by default by `Dapr`; browse the url `http://localhost:9411` to inspect any distributed trace from all microservices

# Product Catalog Service

- REST URL: http://localhost:5002
- gRPC URL: tcp://localhost:50002
- Reference to [how to run](product-catalog/README.md)

# Sale Payment Service

- REST URL: http://localhost:5003

# gRPC

- Reference to [how to generate protobuf](proto/README.md)

# Tools:
- BloomRPC: test gRPC
- vscode.httpclient