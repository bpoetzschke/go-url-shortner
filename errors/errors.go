package errors

import goErr "errors"

var (
	ErrorNotFound = goErr.New("not_found")
)
