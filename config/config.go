package config

// App holds a pointer to the app config create on the Init function call.
// It is there for ease of use on all other needed packages
var App *AppConfig = &AppConfig{}

func Init(appConfigFilename, twitterConfigFilename string) (*AppConfig, error) {
	return nil, nil
}
