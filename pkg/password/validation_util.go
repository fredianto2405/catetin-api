package password

import (
	"errors"
	"regexp"
)

func Validate(password string) error {
	if len(password) < 8 {
		return errors.New("password minimal 8 karakter")
	}

	lowercase := regexp.MustCompile(`[a-z]`)
	uppercase := regexp.MustCompile(`[A-Z]`)
	number := regexp.MustCompile(`[0-9]`)
	special := regexp.MustCompile(`[@#$_&\-+]`)

	switch {
	case !lowercase.MatchString(password):
		return errors.New("password harus mengandung huruf kecil (a-z)")
	case !uppercase.MatchString(password):
		return errors.New("password harus mengandung huruf kapital (A-Z)")
	case !number.MatchString(password):
		return errors.New("password harus mengandung angka (0-9)")
	case !special.MatchString(password):
		return errors.New("password harus mengandung karakter spesial (@#$_&-+)")
	}

	return nil
}
