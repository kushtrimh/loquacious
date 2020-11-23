package apiclient

import (
	"bytes"
	"github.com/kushtrimh/loquacious/auth"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

const (
	mockApiEndpoint  = "http://localhost:8080/api"
	mockAuthEndpoint = "http://localhost:8080/auth"
	accessTokenValue = "Ag35asgdw6326eashasdhaswdyhwqeyqwey"
)

type MockHTTPClient struct{}

func (client *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if strings.HasPrefix(req.Header.Get("Authorization"), "Basic") {
		accessToken := `{"access_token": "` + accessTokenValue + `"}`
		body := bytes.NewReader([]byte(accessToken))
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(body),
			Request:    req,
		}, nil
	}
	return &http.Response{
		StatusCode: 200,
		Request:    req,
	}, nil
}

func TestDoRequest(t *testing.T) {
	apiclient := testAPIClient()
	request, err := http.NewRequest("", "", nil)
	if err != nil {
		t.Error(err)
	}
	response, err := apiclient.Do(request)
	if err != nil {
		t.Errorf("Could not make error for request %v", request)
	}
	if response.StatusCode != 200 {
		t.Errorf("Do(request) does not return %d status code", 200)
	}
}

func TestSuccessfulAuthentication(t *testing.T) {
	apiclient := testAPIClient()
	authData := auth.Auth{Id: "id", Secret: "secret"}
	if err := apiclient.Authenticate(&authData); err != nil {
		t.Error(err)
	}
	if apiclient.token.AccessToken != accessTokenValue {
		t.Errorf("Authenticate(authData) = %s, expected access token value %s",
			apiclient.token.AccessToken, accessTokenValue)
	}
}

func TestSuccessfulGet(t *testing.T) {
	apiclient := testAPIClient()
	apiclient.token = &Token{AccessToken: accessTokenValue}
	response, err := apiclient.Get("", url.Values{})
	if err != nil {
		t.Error(err)
	}
	requestAuthHeader := response.Request.Header.Get("Authorization")
	if requestAuthHeader != "Bearer "+accessTokenValue {
		t.Errorf(`Get("", url.Values{}) sent with bearer token %s, expected bearer token on
				request headers %s`, response.Request.Header.Get("Authorization"), accessTokenValue)
	}
}

func testAPIClient() *APIClient {
	apiclient := New(mockApiEndpoint, mockAuthEndpoint)
	apiclient.client = &MockHTTPClient{}
	return apiclient
}
