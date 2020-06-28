package error

import "fmt"

type VersionNotFound struct {
	ID string
}

func (s *VersionNotFound) Error() string {
	return fmt.Sprintf("version not found with ID: %v", s.ID)
}
