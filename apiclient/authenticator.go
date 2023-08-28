package apiclient

import (
	"encoding/json"
	"strings"

	"github.com/nadavbm/etzqa/pkg/env"
	"github.com/nadavbm/etzqa/pkg/reader"
	"github.com/nadavbm/zlog"
	"gopkg.in/yaml.v2"
)

// Authenticator takes a secret file and authenticate to sql or api server
type Authenticator struct {
	Logger     *zlog.Logger
	reader     *reader.Reader
	SecretFile string
}

// newAuthenticator creates an instance of authenticator
func newAuthenticator(logger *zlog.Logger, secretFile string) *Authenticator {
	return &Authenticator{
		Logger:     logger,
		reader:     reader.NewReader(logger),
		SecretFile: secretFile,
	}
}

// ApiAuth is an api server authentication
type ApiAuth struct {
	// Method is the authentication method, e.g. Bearer or ApiKey
	Method string `json:"method,omitempty" yaml:"method,omitempty"`
	// Token is the authentication token (Bearer token or API key value)
	Token string `json:"token,omitempty" yaml:"token,omitempty"`
}

// GetAPIAuth returns api authentication params from a secret
func (a *Authenticator) GetAPIAuth() (*ApiAuth, error) {
	auth := getAPIAuthFromEnv()
	if auth != nil {
		return auth, nil
	}
	return a.getAPIAuthFromFile()
}

// getAPIAuthFromEnv gets secrets from environment variable (if set)
func getAPIAuthFromEnv() *ApiAuth {
	var auth ApiAuth
	if env.ApiToken != "" && env.ApiAuthMethod != "" {
		auth.Method = env.ApiAuthMethod
		auth.Token = env.ApiToken
		return &auth
	}

	return nil
}

// GetAPIAuth returns api authentication params from a secret
func (a *Authenticator) getAPIAuthFromFile() (*ApiAuth, error) {
	secret, err := a.parseSecret()
	if err != nil {
		return nil, err
	}
	return secret, nil
}

// parseSecret create a secret from json file
func (a *Authenticator) parseSecret() (*ApiAuth, error) {
	bs, err := a.reader.ReadFile(a.SecretFile)
	if err != nil {
		return nil, err
	}

	var s ApiAuth
	switch {
	case strings.HasSuffix(a.SecretFile, ".json"):
		if err := json.Unmarshal(bs, &s); err != nil {
			return nil, err
		}
	case strings.HasSuffix(a.SecretFile, ".yaml"):
		if err := yaml.Unmarshal(bs, &s); err != nil {
			return nil, err
		}
	}

	return &s, nil
}
