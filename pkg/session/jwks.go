package session

import (
	"github.com/zsmatrix62/templ-goat/pkg/config"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/jwks"
)

var (
	jwksClient   *jwks.Client
	_jsonwebKey  *clerk.JSONWebKey
	_clerkConfig *clerk.ClientConfig
)

func GetClerkConfig() *clerk.ClientConfig {
	if _clerkConfig == nil {
		_clerkConfig = &clerk.ClientConfig{}
		_clerkConfig.Key = clerk.String(config.GetString("clerk.secret"))
	}
	return _clerkConfig
}

func GetJWKSClient() *jwks.Client {
	if jwksClient == nil {
		jwksClient = jwks.NewClient(GetClerkConfig())
	}
	return jwksClient
}

func GetJWK() *clerk.JSONWebKey {
	return _jsonwebKey
}

func SetJWK(jwk *clerk.JSONWebKey) {
	_jsonwebKey = jwk
}
