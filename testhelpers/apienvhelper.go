package testhelpers

import (
	"context"
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

func NewTestTools(t *testing.T) (*fiber.App, *api.ApiEnv) {
	apiEnv := NewTestApiEnv(t)
	app := fiber.New()
	app.Route("/users", apiEnv.InitUserRouter())
	app.Route("/books", apiEnv.InitBookRouter())
	app.Route("/auth", apiEnv.InitAuthRouter())
	app.Route("clubs", apiEnv.InitClubRouter())
	return app, apiEnv
}

func WipeDB(e *api.ApiEnv) {
	ctx := context.Background()

	e.Client.User.Delete().Exec(ctx)
	e.Client.Book.Delete().Exec(ctx)
	e.Client.Club.Delete().Exec(ctx)
}
