// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldLastLoginAt holds the string denoting the last_login_at field in the database.
	FieldLastLoginAt = "last_login_at"
	// EdgeClubs holds the string denoting the clubs edge name in mutations.
	EdgeClubs = "clubs"
	// EdgeInvitations holds the string denoting the invitations edge name in mutations.
	EdgeInvitations = "invitations"
	// Table holds the table name of the user in the database.
	Table = "users"
	// ClubsTable is the table that holds the clubs relation/edge. The primary key declared below.
	ClubsTable = "user_clubs"
	// ClubsInverseTable is the table name for the Club entity.
	// It exists in this package in order to avoid circular dependency with the "club" package.
	ClubsInverseTable = "clubs"
	// InvitationsTable is the table that holds the invitations relation/edge.
	InvitationsTable = "invitations"
	// InvitationsInverseTable is the table name for the Invitation entity.
	// It exists in this package in order to avoid circular dependency with the "invitation" package.
	InvitationsInverseTable = "invitations"
	// InvitationsColumn is the table column denoting the invitations relation/edge.
	InvitationsColumn = "invitation_sponsor"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldFirstName,
	FieldLastName,
	FieldPassword,
	FieldEmail,
	FieldCreatedAt,
	FieldLastLoginAt,
}

var (
	// ClubsPrimaryKey and ClubsColumn2 are the table columns denoting the
	// primary key for the clubs relation (M2M).
	ClubsPrimaryKey = []string{"user_id", "club_id"}
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
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultLastLoginAt holds the default value on creation for the "last_login_at" field.
	DefaultLastLoginAt func() time.Time
)
