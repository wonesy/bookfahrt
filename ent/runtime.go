// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/wonesy/bookfahrt/ent/book"
	"github.com/wonesy/bookfahrt/ent/club"
	"github.com/wonesy/bookfahrt/ent/invitation"
	"github.com/wonesy/bookfahrt/ent/schema"
	"github.com/wonesy/bookfahrt/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	bookFields := schema.Book{}.Fields()
	_ = bookFields
	// bookDescTitle is the schema descriptor for title field.
	bookDescTitle := bookFields[0].Descriptor()
	// book.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	book.TitleValidator = bookDescTitle.Validators[0].(func(string) error)
	// bookDescAuthor is the schema descriptor for author field.
	bookDescAuthor := bookFields[1].Descriptor()
	// book.AuthorValidator is a validator for the "author" field. It is called by the builders before save.
	book.AuthorValidator = bookDescAuthor.Validators[0].(func(string) error)
	clubFields := schema.Club{}.Fields()
	_ = clubFields
	// clubDescName is the schema descriptor for name field.
	clubDescName := clubFields[1].Descriptor()
	// club.NameValidator is a validator for the "name" field. It is called by the builders before save.
	club.NameValidator = clubDescName.Validators[0].(func(string) error)
	// clubDescID is the schema descriptor for id field.
	clubDescID := clubFields[0].Descriptor()
	// club.DefaultID holds the default value on creation for the id field.
	club.DefaultID = clubDescID.Default.(func() uuid.UUID)
	invitationFields := schema.Invitation{}.Fields()
	_ = invitationFields
	// invitationDescID is the schema descriptor for id field.
	invitationDescID := invitationFields[0].Descriptor()
	// invitation.DefaultID holds the default value on creation for the id field.
	invitation.DefaultID = invitationDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[3].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[5].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescLastLoginAt is the schema descriptor for last_login_at field.
	userDescLastLoginAt := userFields[6].Descriptor()
	// user.DefaultLastLoginAt holds the default value on creation for the last_login_at field.
	user.DefaultLastLoginAt = userDescLastLoginAt.Default.(func() time.Time)
}
