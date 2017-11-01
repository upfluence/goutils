package stringutil

import (
	"unicode/utf8"

	"golang.org/x/text/encoding/charmap"

	"github.com/upfluence/pkg/log"
)

var defaultDecoder = charmap.ISO8859_1.NewDecoder()

func DecodeToUTF8(s string) string {
	if s == "" {
		return s
	}

	if IsUTF8(s) {
		return s
	}

	s, err := defaultDecoder.String(s)

	if err != nil {
		log.Warning(err.Error())

		return ""
	}

	return s
}

func IsUTF8(s string) bool {
	return utf8.ValidString(s)
}
