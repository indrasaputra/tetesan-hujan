package raindrop

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

// API responsibles to connect to Raindrop.io.
type API struct {
	client *http.Client
}

// NewAPI creates an instance of API.
func NewAPI(token string) *API {
	client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: token,
	}))

	return &API{
		client: client,
	}
}
