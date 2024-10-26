package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	application "github.com/bmena34/wordie/internal/api"
	//"github.com/joho/godotenv"
)

func main() {
	// Uncomment the following lines to load the .env file

	// err := godotenv.Load("../../.env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	// Main function to start the application //

	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx) // change from ":=" to "=" to use .env file
	if err != nil {
		fmt.Println("failed to start application:", err)
	}
}
