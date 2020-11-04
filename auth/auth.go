package auth

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// An Auth struct represent the authentication data
// that is needed to authenticate with an API
type Auth struct {
	Id     string
	Secret string
}

// Base64Encoded encodes the id and secret with base64
// in the format id:secret
func (a *Auth) Base64Encoded() string {
	authentication := a.Id + ":" + a.Secret
	return base64.StdEncoding.EncodeToString([]byte(authentication))
}

// Retrieve the auth data from the give config.
// Error will be returned if the config file does not exist, or it
// cannot be opened
func RetrieveAuthConfig(configFilename string) (*Auth, error) {
	if configFilename == "" {
		errors.New("No config filename specified")
	}
	configPath := authConfigPath(configFilename)
	configExists := AuthConfigExists(configFilename)
	if configExists {
		data, err := ioutil.ReadFile(configPath)
		if err != nil {
			log.Fatal("Could not open auth config file", err)
		}
		return createAuthConfig(data), nil
	}
	return nil, errors.New("Auth config file is corrupted or does not exist")
}

// Create the auth config file with the given id and secret
// The data in the created file will be saved as a JSON object
// An error is returned if the file cannot be created or if any of the parameters
// are missing.
func CreateAuthConfig(id, secret, configFilename string) (*Auth, error) {
	if id == "" || secret == "" || configFilename == "" {
		return nil, errors.New("Id, secret and config filename are required")
	}
	auth := Auth{id, secret}
	authData, err := json.Marshal(auth)
	if err != nil {
		return nil, err
	}
	err = ioutil.WriteFile(authConfigPath(configFilename), authData, 0600)
	if err != nil {
		return nil, err
	}
	return &auth, nil
}

// Check if a file exists for the given filename
func AuthConfigExists(configFilename string) bool {
	_, err := os.Stat(authConfigPath(configFilename))
	if err != nil {
		return !os.IsNotExist(err)
	}
	return true
}

func createAuthConfig(configData []byte) *Auth {
	var auth Auth
	err := json.Unmarshal(configData, &auth)
	if err != nil {
		log.Fatal(err)
	}
	return &auth
}

func authConfigPath(configFilename string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(homeDir, configFilename)
}
