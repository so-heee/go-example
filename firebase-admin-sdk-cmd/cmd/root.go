package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

type Config struct {
	Debug bool

	Firebase struct {
		Credentials string
	}
}

var config Config

const (
	Warning = `WARNING!`
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "firebase-admin-go-cmd",
		Short: "Command line tool using firebase-admin-go",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.HelpFunc()(cmd, args)
			}
		},
	}

	cobra.OnInitialize(initConfig)

	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.firebase-admin-go-cmd.yaml)")
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	cmd.AddCommand(NewCmdGet())
	cmd.AddCommand(NewCmdCreate())
	cmd.AddCommand(NewCmdUpdate())
	cmd.AddCommand(NewCmdDelete())
	return cmd
}

func Execute() {
	cmd := NewCmdRoot()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".firebase-admin-go-cmd.yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println(err)
	}

	if err := viper.Unmarshal(&config); err == nil {
		fmt.Println("firebase credentials file:", config.Firebase.Credentials)
		fmt.Println("debug:", config.Debug)
	} else {
		fmt.Println(err)
	}

}

func getFirebaseClient() (context.Context, *auth.Client) {
	// Get an auth client from the firebase.App
	ctx := context.Background()
	opt := option.WithCredentialsFile(config.Firebase.Credentials)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error initializing app: %v\n", err)
	}
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("Error getting Auth client: %v\n", err)
	}
	return ctx, client
}
