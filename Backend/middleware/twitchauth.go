package middleware

import (
	"bytes"
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

//type gameData struct {
//	name string `json:"name"`
//	clientImage
//}

var data token

func TwitchAuth() error {
	url := "https://id.twitch.tv/oauth2/token?client_id=" + os.Getenv("TWITCH_CLIENT_ID") + "&client_secret=" + os.Getenv("TWITCH_SECRET") + "&grant_type=client_credentials"
	response, err := http.Post(url, "", nil)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	res, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("twitch auth request failed with status: %s, response: %s", response.Status, string(res))
	}

	var datas map[string]interface{}
	if err := json.Unmarshal(res, &datas); err != nil {
		return err
	}

	expiresInVal, ok := datas["expires_in"].(float64)
	if !ok {
		return fmt.Errorf("expires_in not found or not float64 in twitch auth response")
	}
	accessTokenVal, ok := datas["access_token"].(string)
	if !ok {
		return fmt.Errorf("access_token not found or not string in twitch auth response")
	}

	expiresIn := time.Duration(expiresInVal) * time.Second
	data.TokenExpiry = time.Now().Add(expiresIn).Unix()
	data.TokenType = "Bearer"
	data.AccessToken = accessTokenVal
	return nil
}

func GetGames(page string) []byte {
	if data.AccessToken == "" || (data.TokenExpiry) < time.Now().Unix() {
		if err := TwitchAuth(); err != nil {
			fmt.Println("Error in TwitchAuth:", err)
			return nil
		}
	}

	url := "https://api.igdb.com/v4/games"
	body := `
             fields name, cover.image_id, cover.url;
             sort rating desc;
             where rating_count > 100 & version_parent = null;
             limit 20;`

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))

	if err != nil {
		fmt.Println(err)
		return nil
	}

	clientID := os.Getenv("TWITCH_CLIENT_ID")
	acessToken := data.AccessToken
	fmt.Println("acesstoken", acessToken)
	req.Header.Add("Client-ID", clientID)
	req.Header.Add("Authorization", "Bearer "+acessToken)

	client := &http.Client{Timeout: time.Second * 4}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		resBody, _ := io.ReadAll(res.Body)
		fmt.Printf("IGDB API returned status %s: %s\n", res.Status, string(resBody))
		return nil
	}

	read, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	//json.Unmarshal(read, &data)
	fmt.Println(string(read))

	return read
}

func SearchGame(gameName string) []byte {
	url := "https://api.igdb.com/v4/games"

	body := "fields name, cover.url;" + "search " + gameName + ";" + "limit 5"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))

	if err != nil {
		fmt.Println(err)
		return nil
	}
	client := &http.Client{Timeout: time.Second * 5}

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	if res.StatusCode != http.StatusOK {
		fmt.Printf("error ")
		return nil
	}

	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return resBody
}

func GetGameByID(gameId string) []byte {
	url := "https://api.igdb.com/v4/games/"
	body := "fields name, summary, rating, cover.url; where id =" + gameId + ";"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))

	if err != nil {
		fmt.Println(err)
		return nil
	}

	client := &http.Client{Timeout: time.Second * 5}

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Printf("error ")
		return nil
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return resBody
}
