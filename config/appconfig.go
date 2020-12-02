package config

import (
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

func (conf *AppConfig) AddFollowedUser(user string) {
	conf.FollowedUsers = append(conf.FollowedUsers, user)
	err := merge(conf)
	if err != nil {
		log.Fatalf("Could not update configuration when adding user %s, %v",
			user, err)
	}
}

func (conf *AppConfig) FollowedUserExists(user string) bool {
	for _, existingUser := range conf.FollowedUsers {
		if existingUser == user {
			return true
		}
	}
	return false
}

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
