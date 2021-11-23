# northwind-on-dapr
This is a demonstration for how to use Dapr with polyglot programming approach

# Setup env

- Add `.env` file on the root of this repository. That means on https://github.com/vietnam-devs/northwind-on-dapr

```
PRODUCT_HOST=localhost # if on Docker then it should be 0.0.0.0
POSTGRES_USER=northwind
POSTGRES_PASSWORD=<your password>
POSTGRES_DB=northwind_db
```
- Then we can start `tye` as below

```bash
> tye run
```

# Product Catalog 

- REST URL: http://localhost:5002
- gRPC URL: tcp://localhost:50002
- Reference to [how to run](product-catalog/README.md)

# gRPC

- Reference to [how to generate protobuf](proto/README.md)

# Tools:
- BloomRPC: test gRPC
- vscode.httpclient