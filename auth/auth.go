package auth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const authConfig string = ".loquacious-auth.json"

type Auth struct {
	Id     string
	Secret string
}

func RetrieveAuthConfig() (*Auth, error) {
	configPath := authConfigPath()
	configExists := AuthConfigExists()
	if configExists {
		data, err := ioutil.ReadFile(configPath)
		if err != nil {
			log.Fatal("Could not open auth config file", err)
		}
		return createAuthConfig(data), nil
	}
	return nil, errors.New("Auth config file is corrupted or does not exist")
}

func CreateAuthConfig(id, secret string) (*Auth, error) {
	if id == "" || secret == "" {
		return nil, errors.New("Id and secret are required")
	}
	auth := Auth{id, secret}
	authData, err := json.Marshal(auth)
	if err != nil {
		return nil, err
	}
	err = ioutil.WriteFile(authConfigPath(), authData, 0600)
	if err != nil {
		return nil, err
	}
	return &auth, nil
}

func AuthConfigExists() bool {
	_, err := os.Stat(authConfigPath())
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

func authConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(homeDir, authConfig)
}
