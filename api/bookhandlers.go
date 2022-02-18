package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/wonesy/bookfahrt/ent"
	"github.com/wonesy/bookfahrt/ent/book"
)

func (e *ApiEnv) GetAllBooks() ([]*ent.Book, error) {
	return e.Client.Book.Query().All(context.Background())
}

func (e *ApiEnv) GetBookBySlug(slug string) (*ent.Book, error) {
	return e.Client.Book.
		Query().
		Where(book.SlugEQ(slug)).
		Only(context.Background())
}

func (e *ApiEnv) CreateBookHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

func (e *ApiEnv) GetBookHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

func (e *ApiEnv) DeleteBookHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

func (e *ApiEnv) UpdateBookHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
