package model

import (
	"time"
)

type Dog struct {
	ID        uint
	Breed     string
	Variants  *string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type GenericResponse[T any] struct {
	Data   T
	Meta   string
	Error_ string
}

type ListDogsResponseData struct {
	Result  bool
	Message string
	Dogs    []*Dog
}

type CreateDogResponseData struct {
	Result  bool
	Message string
	Dog     *Dog
}
type GetDogByIDResponseData struct {
	Result  bool
	Message string
	Dog     *Dog
}
type DeleteDogByIDResponseData struct {
	Result  bool
	Message string
}

type EditDogByIDResponseData struct {
	Result  bool
	Message string
	Dog     *Dog
}
