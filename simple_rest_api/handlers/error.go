package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// ErrorHandler ...
var ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

// BadRequest ...
type BadRequest struct {
	In    string
	Field string
	Err   error
}

// Error ...
func (br BadRequest) Error() string {
	return fmt.Sprintf("wrong %s field '%s' type: %v", br.In, br.Field, br.Err)
}

// JSONError ...
func JSONError(err error) {
	log.Printf("Error on json writer")
}

// JSONDecodeError ...
type JSONDecodeError struct {
	Err error
}

// Error ...
func (j JSONDecodeError) Error() string {
	return fmt.Sprintf("error on json unmarshal: %v", j.Err.Error())
}

// ErrFieldIsRequired ...
type ErrFieldIsRequired struct {
	Field string
}

// Error ...
func (e ErrFieldIsRequired) Error() string {
	return fmt.Sprintf("'%s' is required", e.Field)
}

// ErrParseField ...
type ErrParseField struct {
	Field string
	In    string
	Err   error
}

// Error ...
func (e ErrParseField) Error() string {
	return fmt.Sprintf("error on parse %s field %q: %v", e.In, e.Field, e.Err.Error())
}

// ErrValidateField ...
type ErrValidateField struct {
	Field string
	In    string
	Msg   string
}

// Error ...
func (e ErrValidateField) Error() string {
	return fmt.Sprintf("error on parse %s field %q: %v", e.In, e.Field, e.Msg)
}
