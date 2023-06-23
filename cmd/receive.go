package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
	"time"
)

var receiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "Get a received SMS",
	Long:  `Get all received SMS from Twilio`,
	Run: func(cmd *cobra.Command, args []string) {

		params := &api.ListMessageParams{}

		params.SetDateSentAfter(time.Unix(fromTimestamp, 0))
		params.SetLimit(limit)

		if watch {
			fmt.Println("🚀 Starting watch mode for new messages")
			params.SetDateSentAfter(time.Now())
			params.SetLimit(1)
			for {
				fmt.Println("⌛ Checking for new messages...")
				receiveMessages(params)
				time.Sleep(time.Second * 10)
			}
		}
		fmt.Println(fmt.Sprintf("🚀 Checking for new messages since given timestamp %d", fromTimestamp))
		receiveMessages(params)
	},
}

func receiveMessages(params *api.ListMessageParams) {
	clientParams := twilio.ClientParams{
		Username: viper.GetString("Account_SID"),
		Password: viper.GetString("Auth_Token"),
	}
	client := twilio.NewRestClientWithParams(clientParams)
	resp, err := client.Api.ListMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for record := range resp {
			if resp[record].Sid != nil {
				fmt.Println(fmt.Sprintf("📧 %s:  %s", *resp[record].DateCreated, *resp[record].Body))
			} else {
				fmt.Println(fmt.Sprintf("📧 %s:  %s", resp[record].DateCreated, resp[record].Body))
			}
		}
	}
}
