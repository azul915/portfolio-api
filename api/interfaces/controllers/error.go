package controllers

// Error は、string型のMessageをDIした struct
type Error struct {
	Message string
}

// NewError は、string型のMessageをDIしたError を返す
func NewError(err error) *Error {
	return &Error{
		Message: err.Error(),
	}
}