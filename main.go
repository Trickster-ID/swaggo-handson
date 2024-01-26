package main

import (
	"context"
	"github.com/joho/godotenv"
	"swaggo-handson/handler"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err.Error())
	}
	handler.MainHttpHandler(context.Background())
}
