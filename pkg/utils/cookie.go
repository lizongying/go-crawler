package utils

import "strings"

var cookieNameSanitizer = strings.NewReplacer("\n", "-", "\r", "-")

func SanitizeCookieName(n string) string {
	return cookieNameSanitizer.Replace(n)
}
