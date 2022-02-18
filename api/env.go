package api

import (
	"github.com/gofiber/fiber/v2/middleware/session"
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
