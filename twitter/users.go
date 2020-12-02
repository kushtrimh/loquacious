package twitter

import (
	"encoding/json"
	"errors"
	"github.com/kushtrimh/loquacious/config"
	"io/ioutil"
	"log"
	"net/url"
)

// User represents the data of a user from Twitter
type User struct {
	Name      string `json:"name"`
	Handle    string `json:"screen_name"`
	Protected bool   `json:"protected"`
}

func (t *Twitter) User(user string) (*User, error) {
	queryParams := url.Values{}
	queryParams.Add("screen_name", user)
	response, err := t.client.Get("/users/show.json", queryParams)
	log.Printf("Queried user %s, response status %d", user, response.StatusCode)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	twitterUser := &User{}
	err = json.Unmarshal(body, twitterUser)
	if err != nil {
		return nil, err
	}
	return twitterUser, nil
}

func (t *Twitter) UserAvailable(user string) bool {
	twitterUser, err := t.User(user)
	if err != nil {
		return false
	}
	return !twitterUser.Protected
}

func (t *Twitter) AddUser(user string) error {
	if config.App.FollowedUserExists(user) {
		return nil
	}
	twitterUser, err := t.User(user)
	if err != nil {
		return err
	}
	if twitterUser.Protected {
		return errors.New("User profile is protected")
	}
	config.App.AddFollowedUser(twitterUser.Handle)
	return nil
}