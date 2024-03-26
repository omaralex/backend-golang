package errors

import (
	"fmt"
)

type InvalidPasswordError struct{}

func (err *InvalidPasswordError) Error() string {
	return fmt.Sprintf("Invalid password")
}

type UserNotFoundError struct{}

func (err *UserNotFoundError) Error() string {
	return fmt.Sprintf("User not found")
}
