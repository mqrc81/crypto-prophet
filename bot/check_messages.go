package bot

import (
	"context"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/adshao/go-binance/v2"

	"github.com/mqrc81/crypto-prophet/database"
	"github.com/mqrc81/crypto-prophet/telegram"
)

const (
	messageRegex = "^((<@!\\d+>).+)+\n" +
		"(?i)SPOT (?i)SIGNAL\n" +
		"[A-Z0-9]+\\/[A-Z]+\n" +
		"ENTER between \\d+.\\d+-\\d+.\\d+\n" +
		"STOP \\d+.\\d+\n" +
		"TARGETS (\\d+(?:.\\d+)?-){5}\\d+(?:.\\d+)?$"

	testMessage = `SPOT SIGNAL
OM/USDT
ENTER between 0.29-0.31
STOP 0.245
TARGETS 0.3179-0.3331-0.3478-0.3767-0.4498-0.5286`
)

var (
	client = binance.NewClient(os.Getenv("BINANCE_API_KEY"), os.Getenv("BINANCE_SECRET_KEY"))
	ctx    = context.Background()
)

func CheckForDiscordMessages() {

	messages, err := database.DB.GetMessages()
	if err != nil {
		log.Printf("error getting discord messages: %v", err)
	}

	if len(messages) < 1 {
		return
	}

	for _, message := range messages {
		signal := parseDiscordMessage(message)
		signal, err = NewBuyOrder(signal)
		if err != nil {
			if err = telegram.SendMessage(telegram.ErrorMSG + "Could not create new buy order with " +
				signal.CryptoBase + "/" + signal.CryptoQuote + " pair."); err != nil {
				log.Printf("error sending telegram message: %v", err)
			}
			continue
		}
		// TODO:
		database.DB.CreateSignal(signal)
	}
}

func parseDiscordMessage(message string) Signal {
	matches, err := regexp.MatchString(message, messageRegex)
	if !matches {
		log.Printf("no regex match: %v", message)
		return Signal{}
	}
	if err != nil {
		if errTG := telegram.SendMessage(telegram.ErrorMSG + err.Error()); errTG != nil {
			log.Printf("error sending telegram message: %v", errTG)
			return Signal{}
		}
	}

	return messageToSignal(message)
}

func messageToSignal(message string) Signal {

	lines := strings.Split(message, "\n")

	cryptoPairs := strings.Split(lines[2], "/")
	if len(cryptoPairs) != 2 {
		if err := telegram.SendMessage(telegram.ErrorMSG + "Could not identify 2 crypto pairs from message."); err != nil {
			log.Printf("error sending telegram message: %v", err)
		}
		return Signal{}
	}

	stopStr := strings.TrimSpace(strings.Trim(lines[4], "STOP "))
	stop, err := strconv.ParseFloat(stopStr, 64)
	if err != nil {
		if err = telegram.SendMessage(telegram.ErrorMSG + "Could not identify stop-loss from message."); err != nil {
			log.Printf("error sending telegram message: %v", err)
		}
		return Signal{}
	}

	var increments []float64
	increments = append(increments, stop, 0) // 0 = placeholder for starting point
	targetsStr := strings.Split(strings.TrimSpace(strings.Trim(lines[5], "TARGETS ")), "-")
	for _, targetStr := range targetsStr {
		target, err := strconv.ParseFloat(targetStr, 64)
		if err != nil {
			if err = telegram.SendMessage(telegram.ErrorMSG + "Could not identify target from message."); err != nil {
				log.Printf("error sending telegram message: %v", err)
			}
			return Signal{}
		}
		increments = append(increments, target)
	}

	return Signal{
		CryptoBase:  cryptoPairs[0],
		CryptoQuote: cryptoPairs[1],
		Increments:  increments,
	}
}
