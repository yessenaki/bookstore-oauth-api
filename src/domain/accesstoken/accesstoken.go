package accesstoken

import "time"

const expirationTime = 24

// AccessToken represents data to access the service
type AccessToken struct {
	AccessToken string `json:"access_token"`
	Expires     int64  `json:"expires"`
	UserID      int    `json:"user_id"`
	ClientID    int    `json:"cliend_id"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now())
}
