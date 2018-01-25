package swgohgg

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var slugRe = regexp.MustCompile("[^a-z0-9]+")

// CharSlug converts the character verbose name into the URL component
func CharSlug(src string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	s, _, _ := transform.String(t, src)
	return strings.Trim(slugRe.ReplaceAllString(strings.ToLower(s), "-"), "-")
}
