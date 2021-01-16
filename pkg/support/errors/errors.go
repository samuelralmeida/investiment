package errors

import (
	"errors"
	"fmt"
	"log"

	"apps/investimento/pkg/ent"
)

var ErrInvalidNota = errors.New("inavlid nota")

func Wrap(msg string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", msg, err)
}

func Log(err error) {
	log.Println(err)
}

func StatusCode(err error) int {
	var nfe *ent.NotFoundError

	// not found
	if errors.As(err, &nfe) {
		return 404
	}

	// bad request
	if errors.Is(err, ErrInvalidNota) {
		return 400
	}

	return 500
}
