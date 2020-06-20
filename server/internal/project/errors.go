package project

import "fmt"

type AlreadyExists struct {
	id int64
}

func (s *AlreadyExists) Error() string {
	return fmt.Sprintf("project already exists with id: %v", s.id)
}

type NotFound struct {
	Id int64
}

func (s *NotFound) Error() string {
	return fmt.Sprintf("project not found with id: %v", s.Id)
}
