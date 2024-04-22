package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printComandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "<Bot User OAuth Token>")
	os.Setenv("SLACK_APP_TOKEN", "<Socket token>")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printComandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples:    []string{"my yob 2024"},
		Handler: func(bc slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("Error")
			}
			age := 2024 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)

		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
