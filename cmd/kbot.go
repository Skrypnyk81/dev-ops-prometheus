/*
Copyright © 2023 Yuriy Skrypnyk <skrypnyk81@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var (
	// TeleToken bot
	TeleToken = os.Getenv("TELE_TOKEN")
)

// kbotCmd represents the kbot command
var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("kbot started %s ", appVersion)
		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})
		if err != nil {
			log.Fatalf("Please check TeleToken env variable %s ", err)
			return
		}
		kbot.Handle("/start", func(c telebot.Context) error {
			return c.Send(fmt.Sprintf("Hello I'm krypto bot version %s\nDigit bitcoin or ethereum", appVersion))
		})
		kbot.Handle(telebot.OnText, func(m telebot.Context) error {
			log.Println(m.Message().Payload, m.Text())
			payload := m.Text()
			var result string
			switch payload {
			case "Bitcoin":
				result = getCrypto("bitcoin")
			case "Ethereum":
				result = getCrypto("ethereum")
			}
			return m.Send(result)
		})

		kbot.Start()
	},
}

type Response struct {
	// Definisci qui la struttura del JSON che vuoi estrarre
	Data struct {
		Symbol    string `json:"symbol"`
		PriceUsd  float64 `json:"priceUsd"`
		Change24H float64 `json:"changePercent24Hr"`
	} `json:"data"`
}

func getCrypto(cry string) string {
	url := "https://api.coincap.io/v2/assets/"
	url += cry
	// Effettua la richiesta GET alla URL
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error for read coincap url %s", err)
	}
	defer response.Body.Close()

	// Decodifica la risposta JSON
	var data Response
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatalf("Error to read JSON %s", err)
	}
	priceUsd, err := strconv.ParseFloat(data.Data.PriceUsd, 64)
	if err != nil {
		log.Fatalf("Error converting priceUsd to float64: %s", err)
	}
	Change24H, err := strconv.ParseFloat(data.Data.Change24H, 64)
	if err != nil {
		log.Fatalf("Error converting priceUsd to float64: %s", err)
	}
	// Stampa i dati estratti dal JSON
	return fmt.Sprintf("Symbol: %s\nPrice in USD: %.2f\nVariation of price in 24H: %.2f%%\n", data.Data.Symbol, priceUsd,
		Change24H)
}

func init() {
	rootCmd.AddCommand(kbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
