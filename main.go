package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mqrc81/crypto-prophet/bot"
)

func main() {
	fmt.Println("Aight let's start over...")

	if os.Getenv("ENV") != "server" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("error loading environment variables: %v", err)
		}
	}

	client, err := bot.AuthorizeTelegram()
	fmt.Println("CLIENT:\n", client, "\n\nERROR:\n", err)

}
