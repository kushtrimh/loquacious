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
	userToAdd    *string = flag.String("add", "", "add user whose tweets you want to count")
	userToRemove *string = flag.String("remove", "", "remove user whose tweets you don't want to count")
	showUsers    *bool   = flag.Bool("users", false, "list all the added users alphabetically")
)

const (
	basedir            string = ".loquacious"
	logFilename        string = basedir + "/loquacious.log"
	authConfigFilename string = basedir + "/lauth.yaml"
	appConfigFilename  string = basedir + "/lapp.yaml"
)

var logFile os.File

func main() {
	logFile, err := os.OpenFile(joinHomeDir(logFilename), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		exit(err)
	}
	log.SetPrefix("[loquacious] ")
	log.SetOutput(logFile)
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

	if *userToAdd != "" {
		t.AddUser(*userToAdd)
		fmt.Printf("User %s added successfully!\n", *userToAdd)
	} else if *userToRemove != "" {
		appConfig.RemoveUser(*userToRemove)
		fmt.Printf("User %s removed successfully!\n", *userToRemove)
	} else {
		// Count all the tweets and display the result
		display(t.TodayTweetCounts())
	}
}

func exit(err error) {
	log.Fatalln(err)
	logFile.Close()
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
		fmt.Printf("%s: %d\n", user, count)
	}
}
