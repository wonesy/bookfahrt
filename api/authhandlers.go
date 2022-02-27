package api

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/wonesy/bookfahrt/auth"
	"github.com/wonesy/bookfahrt/ent"
	"github.com/wonesy/bookfahrt/ent/user"
	"github.com/wonesy/bookfahrt/logging"
)

func (e *ApiEnv) InitAuthRouter() func(fiber.Router) {
	return func(router fiber.Router) {
		router.Post("/login", e.LoginHandler())
		router.Get("/logout", e.LogoutHandler())
		router.Post("/logout", e.LogoutHandler())
		router.Post("/register", e.RegisterHandler())
	}
}

func (e *ApiEnv) LoginHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		creds := new(auth.Credentials)
		if err := c.BodyParser(creds); err != nil {
			return err
		}

		user, err := e.GetUserByUsername(creds.Username)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).
				SendString(fmt.Sprintf("user %s not found", creds.Username))
		}

		if !auth.PasswordMatchesHash(creds.Password, user.Password) {
			return c.Status(fiber.StatusUnauthorized).
				SendString("invalid credentials")
		}

		e.Store.RegisterType(ent.User{})

		sess, err := e.Store.Get(c)
		if err != nil {
			return err
		}

		sess.Set("user", user)

		if err := sess.Save(); err != nil {
			return err
		}

		return c.JSON(user)
	}
}

func (e *ApiEnv) LogoutHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sess, err := e.Store.Get(c)
		if err != nil {
			panic(err)
		}

		return sess.Destroy()
	}
}

func (e *ApiEnv) RegisterHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		userReg := new(auth.UserRegistration)
		if err := c.BodyParser(userReg); err != nil {
			return errors.Wrap(err, "RegisterHandler failed to parse body")
		}

		hashed, err := auth.HashPassword(userReg.Password)
		if err != nil {
			return errors.Wrap(err, "RegisterHandler failed to hash password")
		}

		newUser, err := e.CreateUser(&ent.User{
			Username:  userReg.Username,
			Password:  hashed,
			FirstName: userReg.FirstName,
			LastName:  userReg.LastName,
			Email:     userReg.Email,
		})
		if err != nil {
			return errors.Wrap(err, "RegisterHandler failed to create user")
		}

		// no invitation, return without adding to club
		if userReg.InvitationID == uuid.Nil.String() {
			return c.JSON(newUser)
		}

		parsedInv, err := uuid.Parse(userReg.InvitationID)
		if err != nil {
			logging.ErrorLogger.Println("failed to parse invitation uuid")
			return c.JSON(newUser)
		}

		if err := e.UseAndDeleteInvitation(newUser, parsedInv); err != nil {
			logging.ErrorLogger.Printf("failed to use invitation: %v", err)
			return c.JSON(newUser)
		}

		updatedUser, err := e.Client.User.Query().
			Where(user.UsernameEQ(newUser.Username)).
			WithClubs().
			Only(context.Background())
		if err != nil {
			logging.ErrorLogger.Printf("failed to query user: %v", err)
			return c.JSON(newUser)
		}
		logging.DebugLogger.Printf(
			"added user '%s' to club '%s'",
			updatedUser.Username,
			updatedUser.Edges.Clubs[0].Name,
		)
		return c.JSON(updatedUser)
	}
}
