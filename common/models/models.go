package models

import "github.com/golang-jwt/jwt/v4"

type Credentials struct {
	Password *string `json:"password"`
	Username *string `json:"username"`
}
type CredentialsWebsite struct {
	LatestPingEncrypted *string `json:"latestpingencrypted"`
	Username            *string `json:"username"`
}

type Claims struct {
	Username *string `json:"username"`
	jwt.RegisteredClaims
}

type WebsiteResponse struct {
	LatestWebsitePing *string
}

type LoginResponse struct {
	Login string
}
