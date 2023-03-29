package errors

type ForbiddenError struct{}

type BadRequestError struct{}

func NewForbiddenError() ForbiddenError {
	return ForbiddenError{}
}

func (s ForbiddenError) Error() string {
	return "Access Denied!"
}

func NewBadRequestError() BadRequestError {
	return BadRequestError{}
}
