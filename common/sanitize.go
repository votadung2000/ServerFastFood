package common

import (
	"html"
	"strings"
)

func Sanitize(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}
