package main

import (
	"flag"
	"fmt"
	"github.com/kushtrimh/loquacious/auth"
	"github.com/kushtrimh/loquacious/twitter"
	"os"
	"path/filepath"
)

var (
	clientId     *string = flag.String("client-id", "", "client id of your application")
	clientSecret *string = flag.String("client-secret", "", "client secret of your application")
)

const (
	basedir               string = ".loquacious"
	authConfigFilename    string = ".loquacious-auth.json"
	appConfigFilename     string = basedir + "/.lapp.yaml"
	twitterConfigFilename string = basedir + "/.ltwitter.yaml"
)

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
	t, err := twitter.New(authConfig)
	if err != nil {
		exit(err.Error())
	}
	t.UserTimeline("khajrizi")
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
