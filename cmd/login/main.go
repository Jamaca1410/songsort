package login

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

type User struct {
	Client string
	Secret string
}

type TokenData struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

const loginUrl string = "https://accounts.spotify.com/api/token"

func (login User) GenerateToken() {
	// TODO: See if is necessary to move this to /src/queries

	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": "Basic " + b64.StdEncoding.EncodeToString([]byte(login.Client+":"+login.Secret)),
	}
	body := url.Values{}
	body.Set("grant_type", "client_credentials")

	client := &http.Client{}

	fmt.Printf("%s\n\n", headers["Authorization"])

	req, err := http.NewRequest("POST", loginUrl, bytes.NewBufferString(body.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}
	req.Header.Add("Content-Type", headers["Content-Type"])
	req.Header.Add("Authorization", headers["Authorization"])
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	var Token TokenData
	// fmt.Printf("%s\n", res.Status) DEBUG

	// Print the response body
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}

	if err := json.Unmarshal([]byte(buf.String()), &Token); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}

	// fmt.Println("Response Body:", buf.String()) DEBUG
	fmt.Printf("Client ID: %s\nSecret ID: %s\n", login.Client, login.Secret)
	fmt.Printf("Token: %s\nExpires in: %d\n", Token.AccessToken, Token.ExpiresIn)
}
