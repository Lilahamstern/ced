package error

import "fmt"

type ProjectAlreadyExists struct {
	ID int64
}

func (s *ProjectAlreadyExists) Error() string {
	return fmt.Sprintf("project already exists with ID: %v", s.ID)
}

type ProjectNotFound struct {
	ID int64
}

func (s *ProjectNotFound) Error() string {
	return fmt.Sprintf("project not found with ID: %v", s.ID)
}
