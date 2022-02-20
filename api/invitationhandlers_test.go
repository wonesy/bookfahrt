package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wonesy/bookfahrt/ent"
	"github.com/wonesy/bookfahrt/testhelpers"
)

func TestCreateInvitation(t *testing.T) {
	env := testhelpers.NewTestApiEnv(t)
	defer testhelpers.WipeDB(env)

	// create club
	createdClub, _ := env.CreateClub(&ent.Club{
		Name: "invite-club",
	})

	// create user
	createdUser, _ := env.CreateUser(&ent.User{
		Username: "sponsor",
		Password: "password",
	})

	// create invitation
	createdInv, err := env.CreateInvitation(createdUser, createdClub.ID)
	assert.NoError(t, err)
	assert.NotNil(t, createdInv)

	// validate all edges
	gotInvitation, err := env.GetInvitationByID(createdInv.ID)
	assert.NoError(t, err)
	assert.Equal(t, createdClub.ID, gotInvitation.Edges.Club.ID)
	assert.Equal(t, createdUser.ID, gotInvitation.Edges.Sponsor.ID)
}
