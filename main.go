package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"

	"time"
	"github.com/gospodinzerkalo/bot_covid19/bot"
	"github.com/joho/godotenv"
)

var (
	token = ""
)


func main() {
	// CLI command for starting APP
	app := cli.NewApp()
	app.Commands = cli.Commands{
		&cli.Command{
			Name:   "start",
			Usage:  "start the bot",
			Action: StartBot,
		},
	}
	app.Run(os.Args)

}

func parseEnv() {
	godotenv.Overload(".env")
	token = os.Getenv("TG_TOKEN")
}

func StartBot(d *cli.Context) error {
	parseEnv()
	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		URL:    "",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
	}

	end := bot.NewEndpointsFactory()

	b.Handle("/start",end.Start(b))
	b.Handle(&bot.KeyAll,end.AllCases(b))
	b.Handle(&bot.KeyKz,end.CheckKz(b))
	b.Handle(&bot.KeyByCountry,end.FindByCountry(b))
	fmt.Println("BOT LAUNCHED!!!")
	b.Start()
	return nil
}


