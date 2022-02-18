package api

import (
	"context"
	"errors"
	"fmt"

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

func (e *ApiEnv) CreateBook(book *ent.Book) (*ent.Book, error) {
	if book.Author == "" || book.Title == "" {
		return nil, errors.New("author and title must exist")
	}

	return e.Client.Book.Create().
		SetAuthor(book.Author).
		SetTitle(book.Title).
		SetSlug(GenBookSlug(book)).
		Save(context.Background())
}

func (e *ApiEnv) UpdateBook(b *ent.Book) (*ent.Book, error) {
	fetched, err := e.GetBookBySlug(b.Slug)
	if err != nil {
		return nil, err
	}

	// don't update anything if title and author have not changed
	if fetched.Author == b.Author && fetched.Title == b.Title {
		return fetched, nil
	}

	return fetched.Update().
		SetAuthor(b.Author).
		SetTitle(b.Title).
		SetSlug(GenBookSlug(b)).
		Save(context.Background())
}

func (e *ApiEnv) DeleteBook(slug string) (int, error) {
	return e.Client.Book.Delete().
		Where(book.SlugEQ(slug)).
		Exec(context.Background())
}

func (e *ApiEnv) CreateBookHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var book *ent.Book
		if err := c.BodyParser(book); err != nil {
			return nil
		}
		createdBook, err := e.CreateBook(book)
		if err != nil {
			return err
		}
		return c.JSON(createdBook)
	}
}

func (e *ApiEnv) GetBookHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		slug := c.Params("slug")
		if slug == "" {
			books, err := e.GetAllBooks()
			if err != nil {
				return err
			}
			return c.JSON(books)
		}

		book, err := e.GetBookBySlug(slug)
		if err != nil {
			return err
		}
		return c.JSON(book)
	}
}

func (e *ApiEnv) DeleteBookHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		slug := c.Params("slug")
		numDeleted, err := e.DeleteBook(slug)
		if err != nil {
			return err
		}
		return c.SendString(fmt.Sprintf("Deleted %d book(s)", numDeleted))
	}
}

func (e *ApiEnv) UpdateBookHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var book *ent.Book
		if err := c.BodyParser(book); err != nil {
			return err
		}

		updatedBook, err := e.UpdateBook(book)
		if err != nil {
			return err
		}

		return c.JSON(updatedBook)
	}
}
