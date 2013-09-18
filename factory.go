package errors

import "fmt"

func New(text string, args ...interface{}) error {
	return fmt.Errorf(text, args...)
}
