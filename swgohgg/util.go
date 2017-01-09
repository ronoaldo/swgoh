package swgohgg

import (
	"regexp"
	"strings"
)

var slugRe = regexp.MustCompile("[^a-z0-9]+")

func slug(src string) string {
	return strings.Trim(slugRe.ReplaceAllString(strings.ToLower(src), "-"), "-")
}
