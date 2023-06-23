package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"time"
)

var (
	cfgFile       string
	fromTimestamp int64
	limit         int

	watch bool

	rootCmd = &cobra.Command{
		Use:   "twilio-sms",
		Short: "Twilio SMS is a CLI for watching incoming SMS via Twilio",
		Long:  `A fast and flexible CLI for watching incoming SMS via Twilio built with`,
		Run: func(cmd *cobra.Command, args []string) {
			print("Hello from Twilio SMS")
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.twilio.yaml)")
	receiveCmd.Flags().Int64VarP(&fromTimestamp, "from", "f", time.Now().Unix(), "timestamp to start from")
	receiveCmd.Flags().IntVarP(&limit, "limit", "l", 20, "limit of messages to fetch")
	receiveCmd.Flags().BoolVarP(&watch, "watch", "w", false, "watch for new messages")
	viper.SetDefault("author", "Pascal Giessler <pg@syndikat7.de>")
	viper.SetDefault("license", "MIT")

	rootCmd.AddCommand(receiveCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".twilio")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("🗄 Using config file:", viper.ConfigFileUsed())
	}
}
