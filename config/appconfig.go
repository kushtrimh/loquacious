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
	UserTimelineTweetCount int      `yaml:"userTimelineTweetCount"`
	FollowedUsers          []string `yaml:"followedUsers"`
}

func (conf *AppConfig) String() string {
	return fmt.Sprintf(`
		config: %s,
		userTimelineTweetCount: %d,
		followedUsers: %v`,
		conf.configFilename,
		conf.UserTimelineTweetCount,
		conf.FollowedUsers)
}

// AddFollowedUser adds a user into configuration, and updates
// the configuration file
func (conf *AppConfig) AddFollowedUser(user string) {
	conf.FollowedUsers = append(conf.FollowedUsers, user)
	err := merge(conf)
	if err != nil {
		log.Fatalf("Could not update configuration when adding user %s, %v",
			user, err)
	}
}

// FollowedUserExists checks if the specified user already exists
// in the configuration
func (conf *AppConfig) FollowedUserExists(user string) bool {
	for _, existingUser := range conf.FollowedUsers {
		if existingUser == user {
			return true
		}
	}
	return false
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
