package version

import "fmt"

type NotFound struct {
	Id string
}

func (s *NotFound) Error() string {
	return fmt.Sprintf("version not found with id: %v", s.Id)
}
