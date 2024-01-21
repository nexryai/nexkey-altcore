package system

import "errors"

var (
	ErrorOccurredWhileGeneratingID = errors.New("Error occurred while generating ID")
	InvalidParamsOnServiceCall     = errors.New("Invalid param(s) on service call")
	UnableToCreateDatabaseSession  = errors.New("Unable to create database session")
	CacheNotFound                  = errors.New("No cahce found")
	UnexpectedEmptyString          = errors.New("unexpected empty string")
	UserSuspended                  = errors.New("account has been suspended")
	UserNotFound                   = errors.New("user not found")
	NoteNotFound                   = errors.New("note not found")
)
