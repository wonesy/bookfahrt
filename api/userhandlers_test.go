package api_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
	"github.com/wonesy/bookfahrt/ent"
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

	numUpdated, err3 := apiEnv.UpdateUser(update)
	gotUser, err2 := apiEnv.GetUserByUsername("test-username")

	assert.NoError(t, err3)
	assert.NoError(t, err2)
	assert.Equal(t, 1, numUpdated)

	assert.EqualValues(t, createdUser.CreatedAt, gotUser.CreatedAt)
	assert.Equal(t, "test-username", gotUser.Username)
	assert.Equal(t, "new first name", gotUser.FirstName)
	assert.Equal(t, "test-lastname", gotUser.LastName)
	assert.Equal(t, "test-email@email.com", gotUser.Email)
	assert.Equal(t, "test-password", gotUser.Password)

	// clean up
	num, err := apiEnv.DeleteUser("test-username")
	assert.NoError(t, err)
	assert.Equal(t, 1, num)
}

func TestGetUserHandler(t *testing.T) {
	app := testhelpers.NewTestApp(t)

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

	app := testhelpers.NewTestApp(t)

	defer func() {
		req := httptest.NewRequest("DELETE", "/users/username", nil)
		app.Test(req, 1)
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
