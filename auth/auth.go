package auth

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"os"
)

// An Auth struct represent the authentication data
// that is needed to authenticate with an API
type Auth struct {
	Id     string
	Secret string
}

// Base64Encoded encodes the id and secret with base64
// in the format id:secret
func (a *Auth) Base64Encoded() string {
	authentication := a.Id + ":" + a.Secret
	return base64.StdEncoding.EncodeToString([]byte(authentication))
}

// CreateOrRetrieve creates the auth config data if id and secret are provided,
// or returns the config data if it exists
func CreateOrRetrieve(id, secret, configFilepath string) (*Auth, error) {
	var fl *os.File
	var authConfig *Auth
	var err error

	if id != "" && secret != "" {
		fl, err = os.Create(configFilepath)
		if err != nil {
			return nil, err
		}
		authConfig, err = Create(id, secret, fl)
	} else {
		fl, err = os.Open(configFilepath)
		if err != nil {
			return nil, err
		}
		authConfig, err = Retrieve(fl)
	}
	fl.Close()
	if err != nil {
		return nil, err
	}
	return authConfig, nil
}

// Retrieve retrieves the auth data from the give config.
// Error will be returned if the config file does not exist, or it
// cannot be opened
func Retrieve(reader io.Reader) (*Auth, error) {
	data := make([]byte, 2048)
	_, err := reader.Read(data)
	if err != nil {
		return nil, err
	}
	var authConfig *Auth
	if authConfig, err = createAuthConfig(bytes.Trim(data, "\x00")); err != nil {
		return nil, err
	}
	return authConfig, nil
}

// Create creates the auth config file with the given id and secret
// The data in the created file will be saved as a JSON object
// An error is returned if the file cannot be created or if any of the parameters
// are missing.
func Create(id, secret string, writer io.Writer) (*Auth, error) {
	if id == "" || secret == "" || writer == nil {
		return nil, errors.New("Id, secret and config filename are required")
	}
	auth := Auth{id, secret}
	authData, err := json.Marshal(auth)
	if err != nil {
		return nil, err
	}
	if _, err := writer.Write(authData); err != nil {
		return nil, err
	}
	return &auth, nil
}

func createAuthConfig(configData []byte) (*Auth, error) {
	var auth Auth
	err := json.Unmarshal(configData, &auth)
	if err != nil {
		return nil, err
	}
	return &auth, nil
}
