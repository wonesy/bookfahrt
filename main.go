package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/wonesy/bookfahrt/api"
	"github.com/wonesy/bookfahrt/ent"
)

func initDb() *ent.Client {
	db_host := "bkdb"
	db_port := 5432
	db_user := os.Getenv("POSTGRES_USER")
	db_db := os.Getenv("POSTGRES_DB")
	db_pass := os.Getenv("POSTGRES_PASSWORD")

	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", db_host, db_port, db_user, db_db, db_pass))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema %v", err)
	}

	return client
}

func main() {
	app := fiber.New()
	client := initDb()
	defer client.Close()

	apiEnv := api.NewApiEnv(client)

	app.Route("/users", func(router fiber.Router) {
		router.Post("", apiEnv.CreateUserHandler())
		router.Get("/:username?", apiEnv.GetUserHandler())
		router.Put("/:username", apiEnv.UpdateUserHandler())
		router.Delete("/:username", apiEnv.DeleteUserHandler())
	})

	log.Fatal(app.Listen(":4000"))
}
