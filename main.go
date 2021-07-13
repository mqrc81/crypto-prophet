package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mqrc81/crypto-prophet/bot"
)

func main() {

	if os.Getenv("ENV") != "server" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("error loading environment variables: %v", err)
		}
	}

	tgClient, err := bot.AuthorizeTelegram()
	if err != nil {
		panic(err.Error()) // TODO handle err
	}

	registred, err := tgClient.IsSessionRegistred()
	if err != nil {
		panic(err.Error()) // TODO handle err
	}
	fmt.Println("Registered: ", registred)

}
