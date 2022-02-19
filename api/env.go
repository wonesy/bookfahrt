package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/pkg/errors"
	"github.com/wonesy/bookfahrt/ent"
)

type ApiEnv struct {
	Client *ent.Client
	Store  *session.Store
}

func NewApiEnv(client *ent.Client, store *session.Store) *ApiEnv {
	return &ApiEnv{
		Client: client,
		Store:  store,
	}
}

func (e *ApiEnv) GetSessionUser(c *fiber.Ctx) (*ent.User, error) {
	sess, err := e.Store.Get(c)
	if err != nil {
		return nil, errors.Wrap(err, "session fetching store failed")
	}

	user, ok := sess.Get("user").(ent.User)
	if !ok {
		return nil, errors.New("session invalid user object in store")
	}

	return &user, nil
}
