package errors

type ForbiddenError struct{}

func NewForbiddenError() ForbiddenError {
	return ForbiddenError{}
}

func (s ForbiddenError) Error() string {
	return "Access Denied!"
}
