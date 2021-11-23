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

	a := App{}
	a.Initialize(user, password, dbhost, dbname)

	s := server{}
	s.Initialize(user, password, dbhost, dbname)

	go a.RunApp(fmt.Sprintf("%s:5002", host))
	s.RunGrpcApp(fmt.Sprintf("%s:50002", host))
}
