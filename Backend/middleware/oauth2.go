package middleware

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func HandelOauth2(authCode string) {
	ctx := context.Background()

	conf := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		Scopes:       []string{"https://www.googleapis.com/auth/games"},
		RedirectURL:  os.Getenv("REDIRECT_URL"),
	}

	fmt.Println(conf)
	fmt.Println(ctx)
}
