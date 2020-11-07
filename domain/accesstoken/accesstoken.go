package accesstoken

import (
	"strings"
	"time"

	"github.com/yesseneon/bookstore-utils/errors"
)

const expirationTime = 24

type AccessToken struct {
	AccessToken string `json:"access_token"`
	Expires     int64  `json:"expires"`
	UserID      int    `json:"user_id"`
	ClientID    int    `json:"client_id"`
}

type AccessTokenData struct {
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (data *AccessTokenData) Validate() *errors.RESTError {
	if data.GrantType != "password" && data.GrantType != "client_cridentials" {
		return errors.BadRequest("Invalid grant_type parameter")
	}

	return nil
}

func (at *AccessToken) Validate() *errors.RESTError {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.BadRequest("Invalid access token ID")
	}

	if at.UserID <= 0 {
		return errors.BadRequest("Invalid user ID")
	}

	if at.ClientID <= 0 {
		return errors.BadRequest("Invalid client ID")
	}

	if at.Expires <= 0 {
		return errors.BadRequest("Invalid expiration time")
	}

	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now())
}
