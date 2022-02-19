package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/wonesy/bookfahrt/auth"
	"github.com/wonesy/bookfahrt/ent"
)

func (e *ApiEnv) InitAuthRouter() func(fiber.Router) {
	return func(router fiber.Router) {
		router.Post("/login", e.LoginHandler())
		router.Get("/logout", e.LogoutHandler())
		router.Post("/logout", e.LogoutHandler())
	}
}

func (e *ApiEnv) LoginHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var creds auth.Credentials
		if err := c.BodyParser(&creds); err != nil {
			return err
		}

		user, err := e.GetUserByUsername(creds.Username)
		if err != nil {
			return err
		}

		if !auth.PasswordMatchesHash(creds.Password, user.Password) {
			return errors.New("invalid credentials")
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
