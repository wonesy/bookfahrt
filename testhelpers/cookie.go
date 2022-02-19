package testhelpers

import "net/http"

func GetSessionCookie(cookies []*http.Cookie) *http.Cookie {
	for _, cookie := range cookies {
		if cookie.Name == "session_id" {
			return cookie
		}
	}
	return nil
}
