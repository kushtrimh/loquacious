package main

import (
	"fmt"
	"github.com/kushtrimh/loquacious/auth"
	"github.com/kushtrimh/loquacious/cmd"
	"os"
)

const authConfigFilename string = ".loquacious-auth.json"

func main() {
	var authConfig *auth.Auth
	if *cmd.ClientId != "" && *cmd.ClientSecret != "" {
		var err error
		authConfig, err = auth.CreateAuthConfig(*cmd.ClientId, *cmd.ClientSecret, authConfigFilename)
		if err != nil {
			exit(err.Error())
		}
	} else {
		var err error
		authConfig, err = auth.RetrieveAuthConfig(authConfigFilename)
		if err != nil {
			exit(err.Error())
		}
	}
	fmt.Println(authConfig)
}

func exit(message string) {
	fmt.Fprintf(os.Stderr, message+"\n")
	os.Exit(1)
}
