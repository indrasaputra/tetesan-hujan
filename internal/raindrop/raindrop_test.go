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
		collections := createCollections()
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			b, _ := json.Marshal(collections)
			rw.Write(b)
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

func createCollections() []*entity.Collection {
	return []*entity.Collection{
		&entity.Collection{ID: 1, Name: "Collection-1"},
		&entity.Collection{ID: 2, Name: "Collection-2"},
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
