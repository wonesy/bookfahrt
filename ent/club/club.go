// Code generated by entc, DO NOT EDIT.

package club

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the club type in the database.
	Label = "club"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeMembers holds the string denoting the members edge name in mutations.
	EdgeMembers = "members"
	// EdgeInvitations holds the string denoting the invitations edge name in mutations.
	EdgeInvitations = "invitations"
	// Table holds the table name of the club in the database.
	Table = "clubs"
	// MembersTable is the table that holds the members relation/edge. The primary key declared below.
	MembersTable = "user_clubs"
	// MembersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	MembersInverseTable = "users"
	// InvitationsTable is the table that holds the invitations relation/edge.
	InvitationsTable = "invitations"
	// InvitationsInverseTable is the table name for the Invitation entity.
	// It exists in this package in order to avoid circular dependency with the "invitation" package.
	InvitationsInverseTable = "invitations"
	// InvitationsColumn is the table column denoting the invitations relation/edge.
	InvitationsColumn = "invitation_club"
)

// Columns holds all SQL columns for club fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// MembersPrimaryKey and MembersColumn2 are the table columns denoting the
	// primary key for the members relation (M2M).
	MembersPrimaryKey = []string{"user_id", "club_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
