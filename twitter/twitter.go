package twitter

import (
	"encoding/json"
	"github.com/kushtrimh/loquacious/apiclient"
	"github.com/kushtrimh/loquacious/auth"
	"io/ioutil"
	"log"
	"net/url"
)

const (
	APIEndpoint  string = "https://api.twitter.com/1.1"
	AuthEndpoint string = "https://api.twitter.com/oauth2/token"
)

// Twitter represents a client that can make API calls to the Twitter API 1.1
type Twitter struct {
	client   *apiclient.APIClient
	authData *auth.Auth
}

type Tweet struct {
	Id        string `json:"id_str"`
	CreatedAt string `json:"created_at"`
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

// UserTimeline returns tweets from the timeline of the specified user.
// To retrieve the tweets user should have a public profile
func (t *Twitter) UserTimeline(user string) ([]Tweet, error) {
	queryParams := url.Values{}
	queryParams.Add("screen_name", user)
	queryParams.Add("count", "200")
	response, err := t.client.Get("/statuses/user_timeline.json", queryParams)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	tweets := make([]Tweet, 100)
	json.Unmarshal(body, &tweets)
	log.Printf("Response status code for user timeline %d request. User @%s, tweets returned %d",
		response.StatusCode, user, len(tweets))
	return tweets, nil
}
