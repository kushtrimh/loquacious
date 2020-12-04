package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// AppConfig holds all the configuration used in different parts
// of the application
type AppConfig struct {
	configFilename         string   `yaml:-`
	maxAllowedUsers        int      `yaml:-`
	UserTimelineTweetCount int      `yaml:"userTimelineTweetCount"`
	Users                  []string `yaml:"users"`
}

func (conf *AppConfig) String() string {
	return fmt.Sprintf(`
		config: %s,
		userTimelineTweetCount: %d,
		followedUsers: %v`,
		conf.configFilename,
		conf.UserTimelineTweetCount,
		conf.Users)
}

// AddUser adds a user into configuration, and updates
// the configuration file
func (conf *AppConfig) AddUser(user string) {
	conf.Users = append(conf.Users, user)
	merge(conf)
}

// RemoveUser updates the user slice while removing a user from it,
// if the user exists in the slice, otherwise, the user slice
// will not be changed
func (conf *AppConfig) RemoveUser(user string) {
	index := -1
	users := conf.Users
	for i, u := range users {
		if u == user {
			index = i
			break
		}
	}
	if index == -1 {
		return
	}
	conf.Users = append(users[:index], users[index+1:]...)
	merge(conf)
}

// UserExists checks if the specified user already exists
// in the configuration
func (conf *AppConfig) UserExists(user string) bool {
	for _, existingUser := range conf.Users {
		if existingUser == user {
			return true
		}
	}
	return false
}

// MaximiumUsersReached returns true if no more users
// can be added, and false otherwise
func (conf *AppConfig) MaximiumUsersReached() bool {
	return len(conf.Users) >= conf.maxAllowedUsers
}

// merge updates the configuration file, with the current data
// from the config type
func merge(conf *AppConfig) {
	content, err := yaml.Marshal(conf)
	if err != nil {
		log.Fatalf("Could not marshall configuration %v", err)
	}
	err = ioutil.WriteFile(conf.configFilename, content, 0600)
	if err != nil {
		log.Fatalf("Could not write configuration %v", err)
	}
}

func remove(value string, elements []string) []string {
	index := -1
	for i, el := range elements {
		if el == value {
			index = i
			break
		}
	}
	return append(elements[:index], elements[index+1:]...)
}
