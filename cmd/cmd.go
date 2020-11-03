package cmd

import (
	"flag"
)

var ClientId *string = flag.String("client-id", "", "client id of your application")
var ClientSecret *string = flag.String("client-secret", "", "client secret of your application")

func init() {
	flag.Parse()
}
