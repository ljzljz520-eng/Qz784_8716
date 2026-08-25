package api

import (
	"frontend_go/domain"
	"net/url"
	"strings"
)

func ParseQuery(v url.Values) domain.Query {
	return domain.Query{MeetingID: strings.TrimSpace(v.Get("meeting")), Status: strings.TrimSpace(v.Get("status")), Text: strings.TrimSpace(v.Get("q"))}
}
func StatusCode(err error) int {
	if err == nil {
		return 200
	}
	if strings.Contains(err.Error(), "not found") {
		return 404
	}
	return 400
}
func AllowedMethods(path string) []string {
	if path == "/records" {
		return []string{"GET", "POST"}
	}
	return []string{"GET"}
}
