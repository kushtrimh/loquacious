package main

import (
	"flag"
	"github.com/kushtrimh/loquacious/auth"
	"github.com/kushtrimh/loquacious/config"
	"github.com/kushtrimh/loquacious/twitter"
	"log"
	"os"
	"path/filepath"
)

var (
	clientId     *string = flag.String("client-id", "", "client id of your application")
	clientSecret *string = flag.String("client-secret", "", "client secret of your application")
)

const (
	basedir            string = ".loquacious"
	authConfigFilename string = ".loquacious-auth.json"
	appConfigFilename  string = basedir + "/lapp.yaml"
)

func main() {
	flag.Parse()

	authConfigPath := joinHomeDir(authConfigFilename)
	authConfig, err := auth.CreateOrRetrieve(*clientId, *clientSecret, authConfigPath)
	if err != nil {
		exit(err)
	}

	appConfigPath := joinHomeDir(appConfigFilename)
	appConfig, err := config.Init(appConfigPath)
	if err != nil {
		exit(err)
	}
	log.Println("App config initialized")
	log.Println(appConfig)

	_, err = twitter.New(authConfig)
	if err != nil {
		exit(err)
	}
	// t.UserTimeline("khajrizi")

}

func exit(err error) {
	log.Fatalln(err)
	os.Exit(1)
}

func joinHomeDir(configFilename string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Could not create path for config directory for %s with homedir %s, %v",
			configFilename, homeDir, err)
	}
	return filepath.Join(homeDir, configFilename)
}
