package testhelpers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/wonesy/bookfahrt/api"
	"github.com/wonesy/bookfahrt/ent/enttest"
)

func NewTestApiEnv(t *testing.T) *api.ApiEnv {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	return api.NewApiEnv(client, session.New())
}

func NewTestApp(t *testing.T) *fiber.App {
	apiEnv := NewTestApiEnv(t)
	app := fiber.New()
	app.Route("/users", func(router fiber.Router) {
		router.Post("", apiEnv.CreateUserHandler())
		router.Get("/:username?", apiEnv.GetUserHandler())
		router.Put("/:username", apiEnv.UpdateUserHandler())
		router.Delete("/:username", apiEnv.DeleteUserHandler())
	})
	return app
}
