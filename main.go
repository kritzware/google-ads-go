package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kritzware/google-ads-go/ads"
	"github.com/kritzware/google-ads-go/auth"
)

var (
	command      string
	clientID     string
	clientSecret string
	refreshToken string
	filepath     string
)

var set = flag.NewFlagSet("", flag.ExitOnError)

func init() {
	set.StringVar(&clientID, "client-id", "", "OAuth 2 Client Id")
	set.StringVar(&clientSecret, "client-secret", "", "Oauth 2 Client Secret")
	set.StringVar(&refreshToken, "refresh-token", "", "Google Ads OAuth2 Refresh Token")
	set.StringVar(&filepath, "file", "", "JSON file containing OAuth2 credentials")
}

func main() {
	if len(os.Args) == 1 {
		os.Exit(1)
	}

	set.Parse(os.Args[2:])
	command := os.Args[1]

	switch command {
	case "refresh":
		refresh()
	}
}

func refresh() {
	if filepath != "" {
		refreshFromFile()
		os.Exit(0)
	}

	if clientID == "" && clientSecret == "" && refreshToken == "" {
		fmt.Println("Missing required flags in refresh command:")
		set.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println("\nRefreshing access token with provided credentials..")
	refreshTokenFromCredentials(clientID, clientSecret, refreshToken)
}

func refreshFromFile() {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		fmt.Printf("Credentials file not found at %s\n", filepath)
		os.Exit(1)
	}
	params, err := ads.ReadCredentialsFile(filepath)
	if err != nil {
		panic(err)
	}

	fmt.Println("\nRefreshing access token from provided credentials file..")
	refreshTokenFromCredentials(params.ClientID, params.ClientSecret, params.RefreshToken)
}

func refreshTokenFromCredentials(clientID, clientSecret, refreshToken string) {
	creds := auth.NewCredentials(clientID, clientSecret)
	partialToken := auth.NewPartialToken(refreshToken)

	token, err := auth.RefreshToken(creds, partialToken)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nAccess Token: %s\n", token.AccessToken)
	fmt.Printf("Expires: %s\n", token.Expiry.Format("2006-01-02 15:04:05"))
	fmt.Printf("Valid: %t\n", token.Valid())
}
