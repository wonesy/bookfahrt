// Code generated by entc, DO NOT EDIT.

package invitation

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the invitation type in the database.
	Label = "invitation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeSponsor holds the string denoting the sponsor edge name in mutations.
	EdgeSponsor = "sponsor"
	// EdgeClub holds the string denoting the club edge name in mutations.
	EdgeClub = "club"
	// Table holds the table name of the invitation in the database.
	Table = "invitations"
	// SponsorTable is the table that holds the sponsor relation/edge.
	SponsorTable = "invitations"
	// SponsorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	SponsorInverseTable = "users"
	// SponsorColumn is the table column denoting the sponsor relation/edge.
	SponsorColumn = "invitation_sponsor"
	// ClubTable is the table that holds the club relation/edge.
	ClubTable = "invitations"
	// ClubInverseTable is the table name for the Club entity.
	// It exists in this package in order to avoid circular dependency with the "club" package.
	ClubInverseTable = "clubs"
	// ClubColumn is the table column denoting the club relation/edge.
	ClubColumn = "invitation_club"
)

// Columns holds all SQL columns for invitation fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "invitations"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"invitation_sponsor",
	"invitation_club",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
