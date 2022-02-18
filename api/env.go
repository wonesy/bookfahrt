package api

import (
	"github.com/wonesy/bookfahrt/ent"
)

type ApiEnv struct {
	Client *ent.Client
}

func NewApiEnv(client *ent.Client) *ApiEnv {
	return &ApiEnv{
		Client: client,
	}
}
