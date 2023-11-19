package main

import (
	"flag"
	"fmt"
	"os"
	"songsort/cmd/getartist"
	"songsort/cmd/login"
)

func main() {
	loginCmd := flag.NewFlagSet("login", flag.ExitOnError)
	loginClientId := loginCmd.String("client", " ", "Client ID")
	loginSecretId := loginCmd.String("secret", " ", "Secret client ID")

	artist := flag.String("a", " ", "Get artist info")

	// get artist command
	if len(os.Args) < 2 {
		fmt.Println("Display help message")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "login":
		loginCmd.Parse(os.Args[2:])
		// TODO: Check Authorization login and Access Token login
		// TODO: User env to login as second option
		logAccessToken := login.User{
			Client: *loginClientId,
			Secret: *loginSecretId}
		logAccessToken.GenerateToken()
	default:
		fmt.Println("Display help message")
		os.Exit(1)
	}

	// Ignore This
	kanye := getartist.Artist{Name: *artist}
	kanye.Printer()
}
