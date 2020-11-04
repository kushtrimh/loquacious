package auth

import (
	//"encoding/base64"
	"testing"
)

const configFilename string = ".loquacious-test-auth.json"

func TestCreateAuthConfigEmptyValues(t *testing.T) {
	_, err := CreateAuthConfig("", "", configFilename)
	if err == nil {
		t.Errorf("Create auth config should not be allowed with empty id and secret")
	}

	_, err = CreateAuthConfig("abc", "", configFilename)
	if err == nil {
		t.Errorf("Create auth config should not be allowed with id only")
	}

	_, err = CreateAuthConfig("", "abc", configFilename)
	if err == nil {
		t.Errorf("Create auth config should not be allowed with secret only")
	}

	_, err = CreateAuthConfig("abc", "123", "")
}

func TestCreateAuthConfig(t *testing.T) {
	authentication, err := CreateAuthConfig("myid", "mysecret", configFilename)
	if err != nil {
		t.Errorf("Auth config should have been created successfully")
	}

	if authentication.Id != "myid" || authentication.Secret != "mysecret" {
		t.Errorf("Client id and client secret do not match given arguments")
	}
}
