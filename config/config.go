package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

const maxAllowedUsers int = 200

// App holds a pointer to the app config create on the Init function call.
// It is there for ease of use on all other needed packages
var App *AppConfig = &AppConfig{}

// Init loads the config used for the application.
// If the config does not exist, then it will be created with
// default values, otherwise it will be read from
// the existing config file
func Init(appConfigFilename string) (*AppConfig, error) {
	if _, err := os.Stat(appConfigFilename); os.IsNotExist(err) {
		return createConfig(appConfigFilename)
	}
	content, err := ioutil.ReadFile(appConfigFilename)
	if err != nil {
		return nil, err
	}
	config := &AppConfig{
		configFilename:  appConfigFilename,
		maxAllowedUsers: maxAllowedUsers}
	yaml.Unmarshal(content, config)
	App = config
	return config, nil
}

// createConfig creates the config file in yaml with default values
func createConfig(configFilename string) (*AppConfig, error) {
	dir := filepath.Dir(configFilename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 600)
	}
	fl, err := os.Create(configFilename)
	if err != nil {
		return nil, err
	}
	defer fl.Close()
	config := &AppConfig{
		UserTimelineTweetCount: 200,
		Users:                  []string{},
	}
	content, err := yaml.Marshal(config)
	if err != nil {
		return nil, err
	}
	fl.Write(content)
	App = config
	return config, nil
}
