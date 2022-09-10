package errors

import "fmt"

type WrongPlacementError struct {
	Message string
}

func (e WrongPlacementError) Error() string {
	return fmt.Sprintf("%s", e.Message)
}
