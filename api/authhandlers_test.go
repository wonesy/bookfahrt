package api_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wonesy/bookfahrt/auth"
	"github.com/wonesy/bookfahrt/ent"
	"github.com/wonesy/bookfahrt/testhelpers"
)

func TestLogin(t *testing.T) {
	un := "test"
	pw := "pass"

	app, apienv := testhelpers.NewTestTools(t)

	// create a new user with a new password
	hashed, _ := auth.HashPassword(pw)
	createdUser, err := apienv.CreateUser(&ent.User{
		Username: un,
		Password: hashed,
	})
	assert.NoError(t, err)
	assert.Equal(t, un, createdUser.Username)
	assert.True(t, auth.PasswordMatchesHash(pw, createdUser.Password))

	// try logging in via client
	var buf bytes.Buffer
	e := json.NewEncoder(&buf).Encode(&auth.Credentials{
		Username: un,
		Password: pw,
	})
	assert.NoError(t, e)

	req := httptest.NewRequest("POST", "/auth/login", &buf)
	req.Header.Add("Content-Type", "application/json")

	resp, err2 := app.Test(req, -1)
	assert.NoError(t, err2)
	assert.Equal(t, 200, resp.StatusCode)

	b, _ := io.ReadAll(resp.Body)
	var user *ent.User
	json.Unmarshal(b, &user)

	assert.Equal(t, un, user.Username)

	// verify that the user is saved in the session store
	var sessionID string
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "session_id" {
			sessionID = cookie.Value
		}
	}

	assert.NotEmpty(t, sessionID)
}
