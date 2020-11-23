package config

// AppConfig holds all the configuration used in different parts
// of the application
type AppConfig struct {
	UserTimelineTweetCount int
	FollowedUsers          []TwitterUser
}

// TwitterUser holds information about a specific user from Twitter
type TwitterUser struct {
	Username string
	Id       string
}
