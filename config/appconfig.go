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
	err := merge(conf)
	if err != nil {
		log.Fatalf("Could not update configuration when adding user %s, %v",
			user, err)
	}
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
func merge(conf *AppConfig) error {
	content, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(conf.configFilename, content, 0600)
	if err != nil {
		return err
	}
	return nil
}
