package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/mqrc81/crypto-prophet/bot"
	"github.com/mqrc81/crypto-prophet/database"
)

func main() {

	fmt.Println("Let's get this bread")

	if os.Getenv("ENV") != "heroku" {
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)
		}
	}

	postgresURL := os.Getenv("DATABASE_URL")
	_, err := database.NewDatabase(postgresURL)
	if err != nil {
		log.Fatal(err)
	}

	for {

		go bot.CheckForDiscordMessages()
		go bot.CheckForTradeSignals()

		time.Sleep(3 * time.Minute)
	}

}
