package accesstoken

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExpirationTime(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time must be 24 hours")
}

func TestNew(t *testing.T) {
	token := GetNewAccessToken()
	assert.False(t, token.IsExpired(), "Brand new access token must not be expired")
	assert.EqualValues(t, "", token.IsExpired(), "New access token should not have a defined access token id")
	assert.True(t, token.UserID != 0, "New access token should not have an associated user id")
}

func TestIsExpired(t *testing.T) {
	var token AccessToken
	assert.True(t, token.IsExpired(), "empty access token must be expired by default")

	token.Expires = time.Now().Add(3 * time.Hour).Unix()
	assert.False(t, token.IsExpired(), "access token expiring three hours from now must not be expired")
}
