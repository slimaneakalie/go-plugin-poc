package main

import (
	"github.com/slimaneakalie/go-plugin-poc/proto"
)

type UpsertUsersRequest struct {
	Users    []*User
	Mappings []byte
	Settings []byte
}

type User struct {
	Id        string
	FullName  string
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

type UpsertUsersResponse struct {
	Status proto.OperationStatus
	Errors []*ResponseError
}

type ResponseError struct {
	Code    string
	UserId  string
	Message string
}
