package api_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wonesy/bookfahrt/api"
	"github.com/wonesy/bookfahrt/ent"
)

func TestBookSlug(t *testing.T) {
	book := &ent.Book{
		Author: "Cameron Jones ***",
		Title:  "This'll be good-!",
	}
	re, _ := regexp.Compile("^[0-9]{8}-cameron-jones-thisll-be-good$")
	slug := api.GenBookSlug(book)

	fmt.Println(slug)
	assert.True(t, re.MatchString(slug))
}
