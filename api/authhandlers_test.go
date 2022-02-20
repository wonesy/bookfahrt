package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

	defer func() {
		url := fmt.Sprintf("/users/%s", un)
		req := httptest.NewRequest("DELETE", url, nil)
		app.Test(req, 1)
	}()

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

func TestLogout(t *testing.T) {
	un := "test"
	pw := "pass"

	app, apienv := testhelpers.NewTestTools(t)
	defer testhelpers.WipeDB(apienv)

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

	req = httptest.NewRequest("GET", "/auth/logout", nil)
	resp, err = app.Test(req, 1)
	assert.NoError(t, err)

	// make sure that the cookie has been deleted from the store
	var removedSessionID string
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "session_id" {
			removedSessionID = cookie.Value
		}
	}
	assert.Empty(t, removedSessionID)
}

func TestRegistrationProcess(t *testing.T) {
	app, env := testhelpers.NewTestTools(t)
	defer testhelpers.WipeDB(env)

	// create club
	club, _ := env.CreateClub(&ent.Club{Name: "test club"})

	// create user
	createdUser, _ := env.CreateUser(&ent.User{
		Username: "u",
		Password: "p",
	}, club.ID)

	// create invitation
	inv, _ := env.CreateInvitation(createdUser, club.ID)

	// attempt to register
	registration := &auth.UserRegistration{
		Username:     "new",
		Password:     "newpass",
		FirstName:    "f",
		LastName:     "l",
		Email:        "a@a.a",
		InvitationID: inv.ID.String(),
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(registration)
	req := httptest.NewRequest("POST", "/auth/register", &buf)
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req, -1)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// validate edges
	var newUser ent.User
	b, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(b, &newUser)
	assert.NoError(t, err)
	assert.Equal(t, registration.Username, newUser.Username)
	assert.Equal(t, 1, len(newUser.Edges.Clubs))
	assert.Equal(t, club.Name, newUser.Edges.Clubs[0].Name)

	// validate invitation was deleted
	invs := env.Client.Invitation.Query().AllX(context.Background())
	assert.Equal(t, 0, len(invs))
}
