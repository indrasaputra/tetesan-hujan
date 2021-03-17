package raindrop_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"unsafe"

	"github.com/indrasaputra/tetesan-hujan/entity"
	"github.com/indrasaputra/tetesan-hujan/internal/raindrop"
	"github.com/stretchr/testify/assert"
)

func TestNewAPI(t *testing.T) {
	t.Run("successfully create an instance of API", func(t *testing.T) {
		api := raindrop.NewAPI("http://localhost:8080", "random-token")
		assert.NotNil(t, api)
	})
}

func TestAPI_GetCollections(t *testing.T) {
	t.Run("http call returns error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			panic("collapse")
		}))
		defer server.Close()

		api := createRaindropAPI(server.Client(), server.URL)
		colls, err := api.GetCollections(context.Background())

		assert.NotNil(t, err)
		assert.Empty(t, colls)
	})

	t.Run("response can't be marshalled", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		api := createRaindropAPI(server.Client(), server.URL)
		colls, err := api.GetCollections(context.Background())

		assert.NotNil(t, err)
		assert.Empty(t, colls)
	})

	t.Run("success get collections", func(t *testing.T) {
		wrapper := createCollectionWrapper()
		collections := createCollections()
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Write([]byte(wrapper))
			rw.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		api := createRaindropAPI(server.Client(), server.URL)
		colls, err := api.GetCollections(context.Background())

		assert.Nil(t, err)
		assert.NotEmpty(t, colls)
		assert.Equal(t, collections, colls)
	})
}

func TestAPI_ParseURL(t *testing.T) {
	t.Run("http call returns error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			panic("collapse")
		}))
		defer server.Close()

		api := createRaindropAPI(server.Client(), server.URL)
		url, err := api.ParseURL(context.Background(), "http://url.url")

		assert.NotNil(t, err)
		assert.Nil(t, url)
	})

	t.Run("response can't be marshalled", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		api := createRaindropAPI(server.Client(), server.URL)
		url, err := api.ParseURL(context.Background(), "http://url.url")

		assert.NotNil(t, err)
		assert.Nil(t, url)
	})

	t.Run("success parse url", func(t *testing.T) {
		parsed := createParsedURL()
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			b, _ := json.Marshal(parsed)
			rw.Write(b)
			rw.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		api := createRaindropAPI(server.Client(), server.URL)
		url, err := api.ParseURL(context.Background(), "http://url.url")

		assert.Nil(t, err)
		assert.NotNil(t, url)
		assert.Equal(t, parsed, url)
	})
}

func TestAPI_SaveRaindrop(t *testing.T) {
	t.Run("http call returns error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			panic("collapse")
		}))
		defer server.Close()

		api := createRaindropAPI(server.Client(), server.URL)
		rd := createRaindrop()
		err := api.SaveRaindrop(context.Background(), rd)

		assert.NotNil(t, err)
	})

	t.Run("response code is not between 200 and 299", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()

		api := createRaindropAPI(server.Client(), server.URL)
		rd := createRaindrop()
		err := api.SaveRaindrop(context.Background(), rd)

		assert.NotNil(t, err)
	})

	t.Run("success save raindrop", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		api := createRaindropAPI(server.Client(), server.URL)
		rd := createRaindrop()
		err := api.SaveRaindrop(context.Background(), rd)

		assert.Nil(t, err)
	})
}

func createCollections() []*entity.Collection {
	return []*entity.Collection{
		{ID: 1, Name: "Collection-1"},
		{ID: 2, Name: "Collection-2"},
	}
}

func createCollectionWrapper() string {
	return `{"items":[{"_id":1,"title":"Collection-1"},{"_id":2,"title":"Collection-2"}]}`
}

func createParsedURL() *entity.ParsedURL {
	url := &entity.ParsedURL{
		Error: "",
	}
	url.Item.Title = "http://url.url"
	url.Item.Excerpt = "just a dummy url"
	url.Item.Meta.Canonical = "http://url.url"

	return url
}

func createRaindrop() *entity.Raindrop {
	return &entity.Raindrop{
		Title:        "all URLs in one place",
		Excerpt:      "an application you want to visit",
		Link:         "http://url.url",
		CollectionID: 1,
	}
}

func createRaindropAPI(client *http.Client, url string) *raindrop.API {
	api := &raindrop.API{}
	refClient := reflect.ValueOf(&client).Elem()
	refURL := reflect.ValueOf(&url).Elem()

	apiClient := reflect.ValueOf(api).Elem().Field(0)
	apiClient = reflect.NewAt(apiClient.Type(), unsafe.Pointer(apiClient.UnsafeAddr())).Elem()
	apiClient.Set(refClient)

	apiURL := reflect.ValueOf(api).Elem().Field(1)
	apiURL = reflect.NewAt(apiURL.Type(), unsafe.Pointer(apiURL.UnsafeAddr())).Elem()
	apiURL.Set(refURL)

	return api
}
