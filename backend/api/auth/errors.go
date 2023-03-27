package auth

import "fmt"

type PasswordDoesntMatchError struct {
}

func NewPasswordDoesntMatchError() PasswordDoesntMatchError {
	return PasswordDoesntMatchError{}
}

func (e PasswordDoesntMatchError) Error() string {
	return "AuthError: Password doesn't match!"
}

type InvalidToken struct {
	Reason string
}

func NewInvalidToken(reason string) InvalidToken {
	return InvalidToken{reason}
}

func (e InvalidToken) Error() string {
	return fmt.Sprintf("AuthError: Invalid token: %s", e.Reason)
}
