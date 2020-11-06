package http

import (
	"github.com/kushtrimh/loquacious/auth"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// APIClient holds the structs and metadata needed to make
// requests to the Twitter API
type APIClient struct {
	client       *http.Client
	apiEndpoint  string
	authEndpoint string
	token        *Token
}

// Token represents the bearer token used to authenticate to the API
type Token struct {
	TokenType   string
	AccessToken string
}

// NewAPIClient return pointer to a new ApiClient
// which can be used to make requests to the Twitter API
func NewAPIClient(apiEndpoint, authEndpoint string) *APIClient {
	client := &http.Client{}
	return &APIClient{client, apiEndpoint, authEndpoint, nil}
}

// Get makes a GET request to the specified API endpoint in
// the APIClient
func (api *APIClient) Get(urlPath string) (*http.Response, error) {
	req, err := http.NewRequest("GET", api.apiEndpoint+urlPath, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authentication", "Bearer "+api.token.AccessToken)
	return api.Do(req)
}

// Authenticate accepts an *Auth struct and performs
// authentication to the Twitter API, it sends the id and secret
// as Base64 encoded and sets the Token in the APIClient
// if the request is successful
func (api *APIClient) Authenticate(authData *auth.Auth) error {
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

// Do performs a request to the API endpoint,
// it returns the response if everything was successful or error
func (api *APIClient) Do(request *http.Request) (*http.Response, error) {
	response, err := api.client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
