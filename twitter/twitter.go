package twitter

import (
	"github.com/kushtrimh/loquacious/apiclient"
	"github.com/kushtrimh/loquacious/auth"
)

const (
	APIEndpoint         string = "https://api.twitter.com/1.1"
	AuthEndpoint        string = "https://api.twitter.com/oauth2/token"
	tweetDateTimeLayout string = "Mon Jan 02 15:04:05 MST 2006"
)

// Twitter represents a client that can make API calls to the Twitter API 1.1
type Twitter struct {
	client   *apiclient.APIClient
	authData *auth.Auth
}

// New returns a new Twitter client ready to make API calls.
// The client gets the Bearer token when created, so all the API
// calls are ready to use
func New(authData *auth.Auth) (*Twitter, error) {
	client := apiclient.New(APIEndpoint, AuthEndpoint)
	if err := client.Authenticate(authData); err != nil {
		return nil, err
	}
	return &Twitter{client, authData}, nil
}
