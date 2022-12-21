package data

import (
	"errors"
	"strings"
)

type Qtype uint8

const (
	textType Qtype = iota + 1
	urlType
	emailType
)

func ParseTypes(s string) ([]string, error) {
	s = strings.TrimSpace(strings.ToLower(s))
	return nil, errors.New("hello")
}
