package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
	"github.com/wonesy/bookfahrt/ent"
	"github.com/wonesy/bookfahrt/ent/user"
	"github.com/wonesy/bookfahrt/testhelpers"

	_ "github.com/mattn/go-sqlite3"
)

func TestCreateDeleteUser(t *testing.T) {
	apiEnv := testhelpers.NewTestApiEnv(t)

	newUser := &ent.User{
		Username:  "test-username",
		FirstName: "test-firstname",
		LastName:  "test-lastname",
		Email:     "test-email@email.com",
		Password:  "test-password",
	}

	createdUser, err1 := apiEnv.CreateUser(newUser)
	gotUser, err2 := apiEnv.GetUserByUsername("test-username")

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.Equal(t, "test-username", createdUser.Username)
	assert.Equal(t, "test-firstname", createdUser.FirstName)
	assert.Equal(t, "test-lastname", createdUser.LastName)
	assert.Equal(t, "test-email@email.com", createdUser.Email)
	assert.Equal(t, "test-password", createdUser.Password)

	assert.EqualValues(t, createdUser.CreatedAt, gotUser.CreatedAt)

	// clean up
	num, err := apiEnv.DeleteUser("test-username")
	assert.NoError(t, err)
	assert.Equal(t, 1, num)
}

func TestCreateUpdateDeleteUser(t *testing.T) {
	apiEnv := testhelpers.NewTestApiEnv(t)

	newUser := &ent.User{
		Username:  "test-username",
		FirstName: "test-firstname",
		LastName:  "test-lastname",
		Email:     "test-email@email.com",
		Password:  "test-password",
	}

	createdUser, err1 := apiEnv.CreateUser(newUser)

	assert.NoError(t, err1)
	assert.Equal(t, "test-username", createdUser.Username)
	assert.Equal(t, "test-firstname", createdUser.FirstName)
	assert.Equal(t, "test-lastname", createdUser.LastName)
	assert.Equal(t, "test-email@email.com", createdUser.Email)
	assert.Equal(t, "test-password", createdUser.Password)

	// update
	update := &ent.User{}
	copier.Copy(update, createdUser)
	update.FirstName = "new first name"

	updatedUser, err3 := apiEnv.UpdateUser(update)

	assert.NoError(t, err3)

	assert.EqualValues(t, createdUser.CreatedAt, updatedUser.CreatedAt)
	assert.Equal(t, "test-username", updatedUser.Username)
	assert.Equal(t, "new first name", updatedUser.FirstName)
	assert.Equal(t, "test-lastname", updatedUser.LastName)
	assert.Equal(t, "test-email@email.com", updatedUser.Email)
	assert.Equal(t, "test-password", updatedUser.Password)

	// clean up
	num, err := apiEnv.DeleteUser("test-username")
	assert.NoError(t, err)
	assert.Equal(t, 1, num)
}

func TestGetUserHandler(t *testing.T) {
	app, _ := testhelpers.NewTestTools(t)

	req := httptest.NewRequest("GET", "/users", nil)

	resp, err := app.Test(req, 1)

	b, _ := io.ReadAll(resp.Body)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "[]", string(b))
}

func TestCreateGetUserHandler(t *testing.T) {
	type payload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	app, _ := testhelpers.NewTestTools(t)

	defer func() {
		req := httptest.NewRequest("DELETE", "/users/username", nil)
		app.Test(req, -1)
	}()

	//
	// Create a new user
	//
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(payload{
		Username: "username",
		Password: "password",
	})
	assert.NoError(t, err)

	req := httptest.NewRequest("POST", "/users", &buf)
	req.Header.Add("Content-Type", "application/json")

	resp, err2 := app.Test(req, -1)
	assert.NoError(t, err2)

	b, _ := io.ReadAll(resp.Body)
	var newUser *ent.User
	json.Unmarshal(b, &newUser)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Contains(t, string(b), "\"username\":\"username\"")

	//
	// Get all users
	//
	req = httptest.NewRequest("GET", "/users", nil)
	resp, err = app.Test(req, 1)

	assert.NoError(t, err)
	b, _ = io.ReadAll(resp.Body)
	var allUsers []*ent.User
	json.Unmarshal(b, &allUsers)

	assert.Equal(t, 1, len(allUsers))
	assert.Equal(t, "username", allUsers[0].Username)
	assert.NotEqual(t, "password", allUsers[0].Password)
}

func TestUpdateUserWithClub(t *testing.T) {
	env := testhelpers.NewTestApiEnv(t)

	createdClub, err := env.CreateClub(&ent.Club{
		Name: "Cameron's Club",
	})
	assert.NoError(t, err)

	createdUser, err1 := env.CreateUser(&ent.User{
		Username: "cameron",
		Password: "password",
		Email:    "email@email.com",
	})
	assert.NoError(t, err1)

	updatedUser, err2 := createdUser.Update().
		SetEmail("newemail@email.com").
		AddClubs(createdClub).
		Save(context.Background())
	assert.NoError(t, err2)
	assert.Equal(t, "newemail@email.com", updatedUser.Email)

	gotUser, err3 := env.Client.User.
		Query().
		Where(user.UsernameEQ("cameron")).
		WithClubs().
		Only(context.Background())
	assert.NoError(t, err3)
	assert.Equal(t, createdClub.ID, gotUser.Edges.Clubs[0].ID)
}
