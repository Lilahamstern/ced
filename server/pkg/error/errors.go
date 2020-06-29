package error

type InternalServerError struct {
}

func (s *InternalServerError) Error() string {
	return "internal server error"
}
