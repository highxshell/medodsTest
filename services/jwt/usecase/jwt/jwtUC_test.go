package jwt

import (
	"github.com/go-playground/assert"
	"os"
	"testing"
)

func TestCreateAndValidateAndRefreshTogether(t *testing.T) {
	os.Setenv("SECRET_KEY", "ThisIsSecret")

	guid := "admin"

	jwt := UseCaseImpl{}
	token, err := jwt.CreateToken(guid)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, "", token.AccessToken)
	assert.NotEqual(t, "", token.RefreshToken)

	guid, err = jwt.ValidateToken(token.AccessToken)
	assert.Equal(t, nil, err)
	assert.Equal(t, guid, "admin")

	payload, err := jwt.ValidateRefreshToken(token)

	assert.Equal(t, nil, err)
	assert.Equal(t, guid, payload)
}
