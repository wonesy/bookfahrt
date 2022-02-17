package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/wonesy/bookfahrt/ent"
)

type Env struct {
	db *sql.DB
}

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
	ctx := context.Background()

	client := initDb()
	defer client.Close()

	err := client.User.Create().
		SetUsername("wonesy").
		SetPassword("asdf").
		OnConflict().
		DoNothing().
		Exec(ctx)
	if err != nil {
		fmt.Println("Conflict, ignoring")
	}

	users, err := client.User.Query().All(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		fmt.Println(user.Username)
	}

	app := fiber.New()

	app.Get("/*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Yo %s", c.Params("*"))
		return c.SendString(msg)
	})

	log.Fatal(app.Listen(":4000"))
}
