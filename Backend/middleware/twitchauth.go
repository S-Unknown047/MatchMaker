package middleware

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type token struct {
	AccessToken string `json:"access_token"`
	TokenExpiry int64  `json:"token_expiry"`
	TokenType   string `json:"token_type"`
}

var data token

func TwitchAuth() {
	url := "https://id.twitch.tv/oauth2/token?client_id=" + os.Getenv("TWITCH_CLIENT_ID") + "&client_secret=" + os.Getenv("TWITCH_SECRET") + "&grant_type=client_credentials"
	response, err := http.Post(url, "", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, _ := io.ReadAll(response.Body)
	defer response.Body.Close()
	fmt.Println("response status:", string(res))
	var datas map[string]interface{}

	json.Unmarshal(res, &datas)

	data.TokenExpiry = time.Now().Add(datas["expires_in"].(time.Duration)).Unix()
	data.TokenType = "Bearer"
	data.AccessToken = datas["access_token"].(string)

}

func GetGames() []byte {
	//TwitchAuth()
	if data.AccessToken == "" || data.TokenExpiry < time.Now().Unix() {
		TwitchAuth()
	}

	url := "https://api.igdb.com/v4/games"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	clientID := os.Getenv("TWITCH_CLIENT_ID")
	//req.Header.Add("Content-Type", "application/json")
	acessToken := data.AccessToken
	fmt.Println("acesstoken", acessToken)
	req.Header.Add("Client-ID", clientID)
	req.Header.Add("Authorization", "Bearer "+acessToken)

	client := &http.Client{Timeout: time.Second * 10}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	read, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()
	fmt.Println("response status:", res.Status)
	return read
}
