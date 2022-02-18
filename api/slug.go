package api

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"

	"github.com/wonesy/bookfahrt/ent"
)

func happyString(input string) string {
	re, _ := regexp.Compile("[^a-zA-Z0-9 ]+")
	replaced := re.ReplaceAllString(input, "")
	return strings.ReplaceAll(strings.ToLower(strings.TrimSpace(replaced)), " ", "-")
}

func GenBookSlug(book *ent.Book) string {
	num := rand.Intn(100000000)

	return fmt.Sprintf("%08d-%s-%s",
		num,
		happyString(book.Author),
		happyString(book.Title),
	)
}
