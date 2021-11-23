package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env")

	a := App{}
	a.Initialize(
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	a.RunApp(":5002")
}
