package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env")

	host := os.Getenv("PRODUCT_HOST")

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbhost := os.Getenv("POSTGRES_HOST")
	dbname := os.Getenv("POSTGRES_DB")

	go InitRestServer(user, password, dbhost, dbname, fmt.Sprintf("%s:5002", host))

	s := server{}
	s.Initialize(user, password, dbhost, dbname)

	s.RunGrpcApp(fmt.Sprintf("%s:50002", host))
}
