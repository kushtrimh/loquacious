package auth

import (
	"encoding/base64"
	"fmt"
	"testing"
)

const configFilename string = ".loquacious-test-auth.json"

type MyWriter struct{}

func (writer MyWriter) Write(p []byte) (n int, err error) {
	fmt.Println("Wrote data to auth config file successfully")
	return 0, nil
}

type MyReader struct{}

func (reader MyReader) Read(p []byte) (n int, err error) {
	testDataBytes := []byte("{\"Id\":\"myid\",\"Secret\":\"mysecret\"}")
	for i, val := range testDataBytes {
		p[i] = val
	}
	return 0, nil
}

func TestCreateAuthConfigEmptyValues(t *testing.T) {
	_, err := CreateAuthConfig("", "", MyWriter{})
	if err == nil {
		t.Errorf("Create auth config should not be allowed with empty id and secret")
	}

	_, err = CreateAuthConfig("abc", "", MyWriter{})
	if err == nil {
		t.Errorf("Create auth config should not be allowed with id only")
	}

	_, err = CreateAuthConfig("", "abc", MyWriter{})
	if err == nil {
		t.Errorf("Create auth config should not be allowed with secret only")
	}

	_, err = CreateAuthConfig("abc", "123", MyWriter{})
}

func TestCreateAuthConfig(t *testing.T) {
	authentication, err := CreateAuthConfig("myid", "mysecret", MyWriter{})
	if err != nil {
		t.Errorf("Auth config should have been created successfully")
	}
	if authentication.Id != "myid" || authentication.Secret != "mysecret" {
		t.Errorf("Client id and client secret do not match given arguments")
	}
}

func TestRetrieveAuthConfig(t *testing.T) {
	authentication, err := RetrieveAuthConfig(MyReader{})
	if err != nil {
		t.Errorf("Auth config should have been read successfully")
	}
	if authentication.Id != "myid" || authentication.Secret != "mysecret" {
		t.Errorf("Client id and client secret do not match given arguments")
	}
}

func TestAuthBase64Encoded(t *testing.T) {
	id, secret := "myid", "mysecret"
	encodedIDSecret := base64.StdEncoding.EncodeToString([]byte(id + ":" + secret))
	authentication := Auth{id, secret}

	if encoded := authentication.Base64Encoded(); encodedIDSecret != encoded {
		t.Errorf("Expected encoded value %v, instead value is %v", encodedIDSecret, encoded)
	}

}
