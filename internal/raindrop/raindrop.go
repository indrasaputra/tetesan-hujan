package raindrop

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/indrasaputra/tetesan-hujan/entity"
	"golang.org/x/oauth2"
)

// API responsibles to connect to Raindrop.io.
type API struct {
	client  *http.Client
	baseURL string
}

// NewAPI creates an instance of API.
func NewAPI(baseURL, token string) *API {
	client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: token,
	}))

	return &API{
		client:  client,
		baseURL: baseURL,
	}
}

// GetCollections gets all root collections from raindrop.io
func (a *API) GetCollections(ctx context.Context) ([]*entity.Collection, error) {
	resp, derr := a.client.Get(a.baseURL + "/collections")
	if derr != nil {
		return []*entity.Collection{}, derr
	}
	defer resp.Body.Close()

	var colls []*entity.Collection
	if jerr := json.NewDecoder(resp.Body).Decode(&colls); jerr != nil {
		return []*entity.Collection{}, jerr
	}

	return colls, nil
}

// ParseURL parse an URL to get detailed information from raindrop.io.
func (a *API) ParseURL(ctx context.Context, url string) (*entity.ParsedURL, error) {
	reqURL := fmt.Sprintf("%s/import/url/parse?url=%s", a.baseURL, url)
	req, rerr := http.NewRequest(http.MethodGet, reqURL, nil)
	if rerr != nil {
		return nil, rerr
	}

	resp, derr := a.client.Do(req)
	if derr != nil {
		return nil, derr
	}
	defer resp.Body.Close()

	var parsed entity.ParsedURL
	if jerr := json.NewDecoder(resp.Body).Decode(&parsed); jerr != nil {
		return nil, jerr
	}

	return &parsed, nil
}

// SaveRaindrop saves a raindrop bookmark to specific collection in raindrop.io.
func (a *API) SaveRaindrop(ctx context.Context, raindrop *entity.Raindrop) error {
	body, merr := json.Marshal(raindrop)
	if merr != nil {
		return fmt.Errorf("Marshal serror")
	}

	req, rerr := http.NewRequest(http.MethodPost, a.baseURL+"/raindrop", bytes.NewBuffer(body))
	if rerr != nil {
		return rerr
	}

	resp, derr := a.client.Do(req)
	if derr != nil {
		return derr
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	var tmp interface{}
	json.NewDecoder(resp.Body).Decode(&tmp)
	return fmt.Errorf("[SaveRaindrop] errors: %v", tmp)
}
