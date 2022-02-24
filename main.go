package main

import (
	"context"
	"fmt"
	"log"
	"os"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	_ "github.com/lib/pq"
	"github.com/wonesy/bookfahrt/api"
	"github.com/wonesy/bookfahrt/ent"
	"github.com/wonesy/bookfahrt/logging"

	_ "github.com/wonesy/bookfahrt/docs"
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

// @title BookFahrt API
// @version 0.1.0
// @description Darn tootin' good book readin'
// @termsOfService http://swagger.io/terms/
// @host localhost:4000
// @BasePath /
func main() {
	// init
	app := fiber.New()
	store := session.New()
	client := initDb()
	defer client.Close()
	apiEnv := api.NewApiEnv(client, store)
	logging.InitLoggers()

	// middleware
	app.Get("/docs/*", swagger.HandlerDefault)
	app.Use(recover.New())

	// routers
	app.Route("/auth", apiEnv.InitAuthRouter())
	app.Route("/users", apiEnv.InitUserRouter())
	app.Route("/books", apiEnv.InitBookRouter())
	app.Route("/clubs", apiEnv.InitClubRouter())

	log.Fatal(app.Listen(":4000"))
}
