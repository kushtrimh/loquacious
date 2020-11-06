package http

import (
	"github.com/kushtrimh/loquacious/auth"
	"log"
	"net/http"
	"net/url"
	"path"
)

type ApiClient struct {
	client   *http.Client
	endpoint string
	auth     *auth.Auth
}

func NewApiClient(endpoint string, auth *auth.Auth) *ApiClient {
	client := &http.Client{}
	return &ApiClient{client, endpoint, auth}
}

func (api *ApiClient) Get(path string) (*http.Response, error) {
	req, err := http.NewRequest("GET", api.mergePath(path), nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(api.auth.Id, api.auth.Secret)
	response, err := api.client.Do(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (api *ApiClient) mergePath(urlPath string) string {
	url, err := url.Parse(api.endpoint)
	if err != nil {
		log.Fatal(err)
	}
	return path.Join(url.Path, urlPath)
}
