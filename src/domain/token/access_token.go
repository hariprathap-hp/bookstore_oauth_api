package token

import (
	"fmt"
	"strings"
	"time"

	"github.com/hari/bookstore_users_api/utils/cryptoUtil"
	"github.com/hari/bookstore_users_api/utils/errors"
)

const (
	expiryTime                 = 24
	grantTypePassword          = "password"
	grandTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for password grant type
	Username string `json:"email"`
	Password string `json:"password"`

	// Used for client_credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessTokenRequest) Validate() *errors.RestErr {

	switch at.GrantType {
	case grantTypePassword:
		break

	case grandTypeClientCredentials:
		break

	default:
		return errors.NewHTTPBadRequestError("invalid grant_type parameter")
	}
	return nil
}

type AccessToken struct {
	Access_Token string `json:"access_token"`
	UserId       int64  `json:"user_id"`
	ClientId     int64  `json:"client_id"`
	ExpiryTime   int64  `json:"expiry"`
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.Access_Token = strings.TrimSpace(at.Access_Token)
	if at.Access_Token == "" {
		return errors.NewHTTPBadRequestError("invalid access token id")
	}
	if at.UserId <= 0 {
		return errors.NewHTTPBadRequestError("invalid user id")
	}
	if at.ClientId <= 0 {
		return errors.NewHTTPBadRequestError("invalid client id")
	}
	if at.ExpiryTime <= 0 {
		return errors.NewHTTPBadRequestError("invalid expiration time")
	}
	return nil
}

func GetNewAccessToken(userId int64) AccessToken {
	at := AccessToken{
		UserId:     userId,
		ExpiryTime: time.Now().UTC().Add(time.Hour * expiryTime).Unix(),
	}
	return at
}

func (at AccessToken) isExpired() bool {
	return time.Unix(at.ExpiryTime, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Generate() {
	at.Access_Token = cryptoUtil.GetMD5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.ExpiryTime))
}
