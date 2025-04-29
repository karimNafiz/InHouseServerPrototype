package in_hosue_errors

import (
	"fmt"
)

var EOF *CustomError = &CustomError{
	code:    0,
	message: "end of file",
}

type CustomError struct {
	code    int
	message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("error %d: %s", e.code, e.message)
}
