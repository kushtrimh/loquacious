package main

import (
	"fmt"
	"github.com/kushtrimh/loquacious/auth"
	"github.com/kushtrimh/loquacious/cmd"
	"os"
	"path/filepath"
)

const authConfigFilename string = ".loquacious-auth.json"

func main() {
	var err error

	authConfigPath, err := authConfigHome(authConfigFilename)
	if err != nil {
		exit(err.Error())
	}

	var authConfig *auth.Auth
	var fl *os.File

	if *cmd.ClientId != "" && *cmd.ClientSecret != "" {
		fl, err = os.Create(authConfigPath)
		if err != nil {
			exit(err.Error())
		}
		authConfig, err = auth.CreateAuthConfig(*cmd.ClientId, *cmd.ClientSecret, fl)
	} else {
		fl, err = os.Open(authConfigPath)
		if err != nil {
			exit(err.Error())
		}
		authConfig, err = auth.RetrieveAuthConfig(fl)
	}
	fl.Close()
	if err != nil {
		exit(err.Error())
	}

	fmt.Println(authConfig)
}

func exit(message string) {
	fmt.Fprintf(os.Stderr, message+"\n")
	os.Exit(1)
}

func authConfigHome(configFilename string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, configFilename), nil
}
