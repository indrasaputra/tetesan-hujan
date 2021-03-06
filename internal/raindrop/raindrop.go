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

type collectionWrapper struct {
	Items []struct {
		ID    int64  `json:"_id"`
		Title string `json:"title"`
	} `json:"items"`
}

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

	var wrapper collectionWrapper
	if jerr := json.NewDecoder(resp.Body).Decode(&wrapper); jerr != nil {
		return []*entity.Collection{}, jerr
	}

	return convertWrapperToCollections(wrapper), nil
}

// ParseURL parse an URL to get detailed information from raindrop.io.
func (a *API) ParseURL(ctx context.Context, url string) (*entity.ParsedURL, error) {
	reqURL := fmt.Sprintf("%s/import/url/parse?url=%s", a.baseURL, url)
	resp, derr := a.client.Get(reqURL)
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

	resp, derr := a.client.Post(a.baseURL+"/raindrop", "application/json", bytes.NewBuffer(body))
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

func convertWrapperToCollections(wrapper collectionWrapper) []*entity.Collection {
	var colls []*entity.Collection
	for _, item := range wrapper.Items {
		colls = append(colls, &entity.Collection{ID: item.ID, Name: item.Title})
	}
	return colls
}
