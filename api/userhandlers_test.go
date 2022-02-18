package api_test

import (
	"testing"

	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
	"github.com/wonesy/bookfahrt/api"
	"github.com/wonesy/bookfahrt/ent"
	"github.com/wonesy/bookfahrt/ent/enttest"

	_ "github.com/mattn/go-sqlite3"
)

func TestCreateDeleteUser(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	apiEnv := api.NewApiEnv(client)

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
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	apiEnv := api.NewApiEnv(client)

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
