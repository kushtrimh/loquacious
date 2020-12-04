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

// QueryUser queries a user from twitter and returns
// a representation of it, or an error if something went wrong
func (t *Twitter) QueryUser(user string) (*User, error) {
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
	if response.StatusCode != 200 {
		log.Printf("Could not add user because %s, response status code %d, %s",
			user, response.StatusCode, string(body))
		return nil, errors.New("User is invalid or does not exist")
	}
	twitterUser := &User{}
	err = json.Unmarshal(body, twitterUser)
	if err != nil {
		return nil, err
	}
	return twitterUser, nil
}

// UserAvailable check if the specified user is avaiable
// to be added for querying.
// Only users that exist and that are not 'protected' users,
// are considered available for adding
func (t *Twitter) UserAvailable(user string) bool {
	twitterUser, err := t.QueryUser(user)
	if err != nil {
		return false
	}
	return !twitterUser.Protected
}

// AddUser adds a user to the configuration
// if the user exists and its available for adding
func (t *Twitter) AddUser(user string) error {
	if config.App.UserExists(user) {
		return errors.New("User already exists")
	}
	if config.App.MaximiumUsersReached() {
		return errors.New("Maximum allow users capacity reached")
	}
	twitterUser, err := t.QueryUser(user)
	if err != nil {
		return err
	}
	if twitterUser.Protected {
		return errors.New("User profile is protected")
	}
	config.App.AddUser(twitterUser.Handle)
	return nil
}

// Following returns a slice of handles of the users
// that the specified user is following
func (t *Twitter) Following(user string) ([]string, error) {
	return nil, errors.New("Not implemented yet")
}

// Followers returns a slice of handles of the users
// that follow the specified user
func (t *Twitter) Followers(user string) ([]string, error) {
	return nil, errors.New("Not implemented yet")
}
