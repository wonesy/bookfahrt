package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wonesy/bookfahrt/auth"
	"github.com/wonesy/bookfahrt/ent"
	"github.com/wonesy/bookfahrt/ent/user"
	"github.com/wonesy/bookfahrt/testhelpers"
)

func TestLoginAndJoinClubHandlers(t *testing.T) {
	app, env := testhelpers.NewTestTools(t)

	// create a club
	createdClub, err := env.CreateClub(&ent.Club{Name: "daclub"})
	assert.NoError(t, err)

	// create a user, specifically without linking the club
	hashed, _ := auth.HashPassword("password")
	createdUser, err := env.CreateUser(&ent.User{
		Username: "cameron",
		Password: hashed,
	})
	_ = createdUser
	assert.NoError(t, err)

	// login, set the session token
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(&auth.Credentials{
		Username: "cameron",
		Password: "password",
	})
	loginReqest := httptest.NewRequest("POST", "/auth/login", &buf)
	loginReqest.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(loginReqest, -1)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// join a book club
	url := fmt.Sprintf("/clubs/%s/join", createdClub.ID)
	joinRequest := httptest.NewRequest("POST", url, nil)
	joinRequest.AddCookie(testhelpers.GetSessionCookie(resp.Cookies()))
	resp, err = app.Test(joinRequest, -1)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// verify the edge
	cameron, err := env.Client.User.Query().Where(user.UsernameEQ("cameron")).WithClubs().Only(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 1, len(cameron.Edges.Clubs))
	assert.Equal(t, createdClub.ID, cameron.Edges.Clubs[0].ID)
	assert.Equal(t, createdClub.Name, cameron.Edges.Clubs[0].Name)
}
