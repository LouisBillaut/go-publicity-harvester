package entity

import "fmt"

type InvalidArgument struct {
	ErrorString string
}

type InvalidType struct {}

type AlreadyExist struct {
	Value string
}

type NotFound struct {
	Value string
}

func (e InvalidArgument) Error() string {
	return fmt.Sprintf("invalid argument, %s", e.ErrorString)
}

func (e InvalidType) Error() string {
	return "this type doesn't exist"
}

func (e NotFound) Error() string {
	return fmt.Sprintf("%s not found", e.Value)
}

func (e AlreadyExist) Error() string {
	return fmt.Sprintf("%s already exist", e.Value)
}

