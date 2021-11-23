# Get starting

- Add `.env` file on the root of this repository. That means on https://github.com/vietnam-devs/northwind-on-dapr

```
POSTGRES_USER=northwind
POSTGRES_PASSWORD=<your password>
POSTGRES_DB=northwind_db
```

- Run

```bash
> go mod download
```

- Then 

```bash
> tye run --tags inf
```

```bash
> go build -o product.bin
> go ./product.bin
```

Ready to go!!!

Notes:
- Install golang on your machine: `1.17.3`
- Install vscode extensions:
  - golang.go
  - doggy8088.go-extension-pack