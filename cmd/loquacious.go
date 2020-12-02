package main

import (
	"flag"
	"fmt"
	"github.com/kushtrimh/loquacious/auth"
	"github.com/kushtrimh/loquacious/config"
	"github.com/kushtrimh/loquacious/twitter"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	clientId     *string = flag.String("client-id", "", "client id of your application")
	clientSecret *string = flag.String("client-secret", "", "client secret of your application")
	userToAdd    *string = flag.String("add-user", "", "add user whose tweets you want to count")
)

const (
	basedir            string = ".loquacious"
	authConfigFilename string = basedir + "/lauth.yaml"
	appConfigFilename  string = basedir + "/lapp.yaml"
)

func main() {
	log.SetPrefix("[loquacious] ")
	flag.Parse()

	appConfigPath := joinHomeDir(appConfigFilename)
	appConfig, err := config.Init(appConfigPath)
	if err != nil {
		exit(err)
	}
	log.Printf("App config initialized %v\n", appConfig)

	authConfigPath := joinHomeDir(authConfigFilename)
	authConfig, err := auth.CreateOrRetrieve(*clientId, *clientSecret, authConfigPath)
	if err != nil {
		exit(err)
	}
	log.Println("Auth config initialized")

	t, err := twitter.New(authConfig)
	if err != nil {
		exit(err)
	}
	log.Println("Twitter API client initialized")

	userToAdd := *userToAdd
	if userToAdd != "" {
		if err := t.AddUser(userToAdd); err != nil {
			fmt.Printf("Could not add user %s, %v", userToAdd, err)
			os.Exit(1)
		}
		fmt.Printf("User %s added successfully!\n", userToAdd)
	} else {
		// Count all the tweets and display the result
		display(t.TodayTweetCounts())
	}
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

func display(tweetCounts map[string]int) {
	now := time.Now()
	fmt.Printf("Today is %s", now.Format("02/01/2006 (Mon)\n"))
	for user, count := range tweetCounts {
		fmt.Printf("%s: %d tweets\n", user, count)
	}
}
