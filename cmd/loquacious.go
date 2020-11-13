package main

import (
	"flag"
	"fmt"
	"github.com/kushtrimh/loquacious/auth"
	"os"
	"path/filepath"
)

var (
	clientId     *string = flag.String("client-id", "", "client id of your application")
	clientSecret *string = flag.String("client-secret", "", "client secret of your application")
)

const authConfigFilename string = ".loquacious-auth.json"

func main() {
	flag.Parse()

	authConfigPath, err := authConfigHome(authConfigFilename)
	if err != nil {
		exit(err.Error())
	}

	authConfig, err := auth.CreateOrRetrieve(*clientId, *clientSecret, authConfigPath)
	if err != nil {
		exit(err.Error())
	}
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
