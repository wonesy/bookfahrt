package api_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wonesy/bookfahrt/ent"
	"github.com/wonesy/bookfahrt/testhelpers"
)

func TestCreateDeleteBook(t *testing.T) {
	apiEnv := testhelpers.NewTestApiEnv(t)
	defer testhelpers.WipeDB(apiEnv)

	book := &ent.Book{
		Title:  "test title 123 ^& something",
		Author: "test author",
	}
	createdBook, err1 := apiEnv.CreateBook(book)
	num, err2 := apiEnv.DeleteBook(createdBook.Slug)

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.Equal(t, 1, num)
	assert.Equal(t, createdBook.Author, book.Author)
	assert.Equal(t, createdBook.Title, book.Title)
	assert.Regexp(t, regexp.MustCompile("^[0-9]{8}-test-author-test-title-123--something$"), createdBook.Slug)
}

func TestCreateUpdateDeleteBook(t *testing.T) {
	apiEnv := testhelpers.NewTestApiEnv(t)
	book := &ent.Book{
		Title:  "test title 123 ^& something",
		Author: "test author",
	}

	createdBook, err1 := apiEnv.CreateBook(book)     // create the book
	sameBook, err2 := apiEnv.UpdateBook(createdBook) // attempt to update without changing

	toUpdate := &ent.Book{
		Title:  "new title",
		Author: "new author",
		Slug:   sameBook.Slug,
	}

	updatedBook, err3 := apiEnv.UpdateBook(toUpdate) // real update
	num, err4 := apiEnv.DeleteBook(updatedBook.Slug) // delete

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.NoError(t, err3)
	assert.NoError(t, err4)

	// test createdBook
	assert.Equal(t, createdBook.Author, book.Author)
	assert.Equal(t, createdBook.Title, book.Title)
	assert.Regexp(t, regexp.MustCompile("^[0-9]{8}-test-author-test-title-123--something$"), createdBook.Slug)

	// test sameBook
	assert.Equal(t, createdBook.Author, sameBook.Author)
	assert.Equal(t, createdBook.Title, sameBook.Title)
	assert.Equal(t, createdBook.ID, sameBook.ID)
	assert.Equal(t, createdBook.Slug, sameBook.Slug)

	// test updatedBook
	assert.Equal(t, toUpdate.Author, updatedBook.Author)
	assert.Equal(t, toUpdate.Title, updatedBook.Title)
	assert.Regexp(t, regexp.MustCompile("^[0-9]{8}-new-author-new-title$"), updatedBook.Slug)
	assert.Equal(t, createdBook.ID, updatedBook.ID)

	assert.Equal(t, 1, num)
}

func TestCreateGetBook(t *testing.T) {
	apiEnv := testhelpers.NewTestApiEnv(t)
	defer testhelpers.WipeDB(apiEnv)

	// assert we start out fresh
	books, err := apiEnv.GetAllBooks()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(books))

	createdBook1, err1 := apiEnv.CreateBook(&ent.Book{
		Title:  "t",
		Author: "a",
	})
	createdBook2, err2 := apiEnv.CreateBook(&ent.Book{
		Title:  "t",
		Author: "a",
	})
	createdBook3, err3 := apiEnv.CreateBook(&ent.Book{
		Title:  "tt",
		Author: "a",
	})

	assert.NoError(t, err1)
	assert.Error(t, err2) // fails because unique index violation
	assert.NoError(t, err3)

	// test createdBook1
	assert.Equal(t, createdBook1.Author, "a")
	assert.Equal(t, createdBook1.Title, "t")
	assert.Regexp(t, regexp.MustCompile("^[0-9]{8}-a-t$"), createdBook1.Slug)

	// test createdBook2
	assert.Nil(t, createdBook2)

	// test createdBook3
	assert.Equal(t, createdBook3.Author, "a")
	assert.Equal(t, createdBook3.Title, "tt")
	assert.Regexp(t, regexp.MustCompile("^[0-9]{8}-a-tt$"), createdBook3.Slug)

	books, err = apiEnv.GetAllBooks()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(books))
}
