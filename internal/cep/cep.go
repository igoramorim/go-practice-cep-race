package cep

import (
	"errors"
	"strings"
)

var (
	ErrInvalidLen  = errors.New("invalid length. CEP is 8 digits long")
	ErrInvalidChar = errors.New("invalid characters. only digits are acceptable")
)

func New(cep string) (CEP, error) {
	if len(cep) != 8 {
		return CEP{}, ErrInvalidLen
	}

	if !isOnlyDigits(cep) {
		return CEP{}, ErrInvalidChar
	}

	prefix := getPrefix(cep)
	suffix := getSuffix(cep)

	return CEP{
		prefix: prefix,
		suffix: suffix,
	}, nil
}

func isOnlyDigits(cep string) bool {
	isNotDigit := func(c rune) bool {
		return c < '0' || c > '9'
	}
	return strings.IndexFunc(cep, isNotDigit) == -1
}

func getPrefix(cep string) string {
	return cep[:5]
}

func getSuffix(cep string) string {
	return cep[5:]
}

type CEP struct {
	prefix string
	suffix string
}

func (cep CEP) Prefix() string {
	return cep.prefix
}

func (cep CEP) Suffix() string {
	return cep.suffix
}
