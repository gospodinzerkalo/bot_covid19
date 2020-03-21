package main

import (
	"github.com/urfave/cli"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"time"
	"./bot"
)

func main() {
	app := cli.NewApp()
	// app.Flags = flags
	app.Commands = cli.Commands{
	&cli.Command{
		Name:   "start",
		Usage:  "start the bot",
		Action: StartBot,
	},
	}
	app.Run(os.Args)
}

func StartBot(d *cli.Context) error {

	b, err := tb.NewBot(tb.Settings{
		Token:  "731713099:AAHjpa_A8Ewv6CURNqfVqfF7AZ4eVKfwqhM",
		URL:    "",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return err
	}

	end := bot.NewEndpointsFactory()

	b.Handle("/start",end.Start(b))
	b.Handle(&bot.KeyAll,end.AllCases(b))
	b.Handle(&bot.KeyKz,end.CheckKz(b))
	b.Handle(&bot.KeyByCountry,end.FindByCountry(b))
	b.Start()

	return nil
}