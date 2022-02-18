package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/wonesy/bookfahrt/auth"
	"github.com/wonesy/bookfahrt/ent"
	"github.com/wonesy/bookfahrt/ent/user"
)

func (e *ApiEnv) GetAllUsers() ([]*ent.User, error) {
	return e.Client.User.Query().All(context.Background())
}

func (e *ApiEnv) GetUserByUsername(username string) (*ent.User, error) {
	return e.Client.User.
		Query().
		Where(user.UsernameEQ(username)).
		Only(context.Background())
}

func (e *ApiEnv) CreateUser(user *ent.User) (*ent.User, error) {
	return e.Client.User.Create().
		SetUsername(user.Username).
		SetFirstName(user.FirstName).
		SetLastName(user.LastName).
		SetEmail(user.Email).
		SetPassword(user.Password).
		Save(context.Background())
}

func (e *ApiEnv) UpdateUser(u *ent.User) (int, error) {
	return e.Client.User.Update().
		Where(user.UsernameEQ(u.Username)).
		SetEmail(u.Email).
		SetFirstName(u.FirstName).
		SetLastName(u.LastName).
		Save(context.Background())
}

func (e *ApiEnv) DeleteUser(username string) (int, error) {
	return e.Client.User.Delete().
		Where(user.UsernameEQ(username)).
		Exec(context.Background())
}

func (e *ApiEnv) GetUserHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		username := c.Params("username")

		if username == "" {
			users, err := e.GetAllUsers()
			if err != nil {
				return err
			}
			return c.JSON(users)
		}

		user, err := e.GetUserByUsername(username)
		if err != nil {
			return err
		}
		return c.JSON(user)
	}
}

func (e *ApiEnv) CreateUserHandler() func(c *fiber.Ctx) error {
	type bodyPassword struct {
		Password string `json:"password"`
	}

	return func(c *fiber.Ctx) error {
		user := new(ent.User)
		pw := new(bodyPassword)

		if err := c.BodyParser(pw); err != nil {
			return err
		}

		hashedPass, err := auth.HashPassword(pw.Password)
		if err != nil {
			return err
		}

		// because password is a sensitive field, it isn't encoded
		// automatically. It must be manually injected into the obj
		if err := c.BodyParser(user); err != nil {
			return err
		}

		user.Password = hashedPass

		newUser, err := e.CreateUser(user)
		if err != nil {
			return err
		}

		return c.JSON(newUser)
	}
}

func (e *ApiEnv) UpdateUserHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user := new(ent.User)

		if err := c.BodyParser(user); err != nil {
			return err
		}

		if c.Params("username") != user.Username {
			return errors.New("usernames must match")
		}

		numRecordsUpdated, err := e.UpdateUser(user)
		if err != nil {
			return err
		}

		return c.SendString(fmt.Sprintf("Updated %d record(s)", numRecordsUpdated))
	}
}

func (e *ApiEnv) DeleteUserHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		username := c.Params("username")

		numRecordsDeleted, err := e.DeleteUser(username)
		if err != nil {
			return err
		}

		return c.SendString(fmt.Sprintf("Deleted %d record(s)", numRecordsDeleted))
	}
}
