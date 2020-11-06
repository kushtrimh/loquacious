package http

import (
	"github.com/kushtrimh/loquacious/auth"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type ApiClient struct {
	client       *http.Client
	apiEndpoint  string
	authEndpoint string
	token        *Token
}

type Token struct {
	TokenType   string
	AccessToken string
}

func NewApiClient(apiEndpoint, authEndpoint string) *ApiClient {
	client := &http.Client{}
	return &ApiClient{client, apiEndpoint, authEndpoint, nil}
}

func (api *ApiClient) Get(urlPath string) (*http.Response, error) {
	req, err := http.NewRequest("GET", api.apiEndpoint+urlPath, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authentication", "Bearer "+api.token.AccessToken)
	return api.Do(req)
}

func (api *ApiClient) Authenticate(authData *auth.Auth) error {
	data := url.Values{}
	data.Add("grant_type", "client_credentials")
	req, err := http.NewRequest("POST", api.authEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(authData.Id, authData.Secret)
	response, err := api.Do(req)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return err
	}
	log.Println(body)
	// TODO: set access token to api client
	return nil
}

func (api *ApiClient) Do(request *http.Request) (*http.Response, error) {
	response, err := api.client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
