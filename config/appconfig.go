package config

// AppConfig holds all the configuration used in different parts
// of the application
type AppConfig struct {
	UserTimelineTweetCount int           `yaml:"userTimelineTweetCount"`
	FollowedUsers          []TwitterUser `yaml:"followedUsers"`
}

// TwitterUser holds information about a specific user from Twitter
type TwitterUser struct {
	Username string `yaml:"username"`
	Id       string `yaml:"id"`
}

// FollowedUsernames returns a []string slice of username
// for all the followed users
func (conf *AppConfig) FollowedUsernames() []string {
	usernames := make([]string, len(conf.FollowedUsers))
	for _, user := range conf.FollowedUsers {
		usernames = append(usernames, user.Username)
	}
	return usernames
}