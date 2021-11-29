# Setup to generate gRPC stubs on Windows

## Step 1

- Download: https://github.com/protocolbuffers/protobuf/releases
- Put to your Windows machine, then set path to run `protoc` 

## Step 2

- Stand in `product-catalog/` folder, then run

```bash
> go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Step 3

At the root of project, that mean is at `https://github.com/vietnam-devs/northwind-on-dapr`

```bash
> protoc protos/product_api.proto -I. --go_out=./product-catalog/ --go_opt=paths=source_relative --go-grpc_out=./product-catalog/ --go-grpc_opt=paths=source_relative
```