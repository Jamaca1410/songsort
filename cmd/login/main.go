package login

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	// TODO: Unmarshal response and convert to struct
	// TODO: Fix how I do the query
	// TODO: Figure how to use res.Body un ma Unmarshal operation
	body := []byte(`"grant_type=client_credentials&client_id=your-client-id&client_secret=your-client-secret"`)

	res, err := http.Post(loginUrl, "application/x-www-form-urlencoded", bytes.NewBuffer(body))

	defer res.Body.Close()

	var Token TokenData
	// TODO: Fix error output
	data, err := json.Marshal(res.Body)
	if err != nil {
		os.Exit(1)
	}

	if err := json.Unmarshal(data, &Token); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}

	if err != nil {
		fmt.Printf("Cliend ID: %s\nSecret ID: %s\nResponse: %s", login.Client, login.Secret, res.AccessToken)
	}

}
