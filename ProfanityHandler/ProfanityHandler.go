package ProfanityHandler

import (
	goaway "github.com/TwiN/go-away"
)

type ProfanityHandler interface {
	RemoveProfanity(text *string)
}

func RemoveProfanity(text *string) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	*text = goaway.Censor(*text)
	return
}
