package apiconfig

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

func AuthConfig() (*Auth, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	authConfigPath := filepath.Join(homeDir, authConfig)
	if _, err := os.Stat(authConfigPath); !os.IsNotExist(err) {
		data, err := ioutil.ReadFile(authConfigPath)
		if err != nil {
			log.Fatal("Could not open auth config file", err)
		}
		return createAuthConfig(data), nil
	}
	return nil, errors.New("Auth config file is corrupted or does not exist")
}

func createAuthConfig(configData []byte) *Auth {
	var auth Auth
	err := json.Unmarshal(configData, &auth)
	if err != nil {
		log.Fatal(err)
	}
	return &auth
}
