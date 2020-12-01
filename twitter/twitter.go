package twitter

import (
	"encoding/json"
	"github.com/kushtrimh/loquacious/apiclient"
	"github.com/kushtrimh/loquacious/auth"
	"github.com/kushtrimh/loquacious/config"
	"io/ioutil"
	"log"
	"net/url"
	"strconv"
	"time"
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
	queryParams.Add("count", strconv.Itoa(config.App.UserTimelineTweetCount))
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

// TodayTweetCount returns the number of tweets made today
// for the specified user
func (t *Twitter) TodayTweetCount(user string) (int, error) {
	tweets, err := t.UserTimeline(user)
	if err != nil {
		return 0, err
	}
	count := 0
	thisDay, thisMonth, thisYear := time.Now().Date()
	for _, tweet := range tweets {
		tweetTime, err := time.Parse(tweetDateTimeLayout, tweet.CreatedAt)
		if err != nil {
			log.Println(err)
			continue
		}
		day, month, year := tweetTime.Date()
		if day == thisDay && month == thisMonth && year == thisYear {
			count++
		}
	}
	return count, nil
}
