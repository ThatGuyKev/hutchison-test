package routes

import (
	"hutchison-test/common"
	"hutchison-test/handler"
	"net/http"
)

var Routes = []common.Route{
	{
		Name:        "CreateDog",
		Method:      http.MethodPost,
		Pattern:     "/api/dogs",
		HandlerFunc: handler.CreateDogHandler, // TODO: add handler func
	},
	{
		Name:        "ListDogs",
		Method:      http.MethodGet,
		Pattern:     "/api/dogs",
		HandlerFunc: handler.ListDogsHandler, // TODO: add handler func
	},
	{
		Name:        "GetDogByID",
		Method:      http.MethodGet,
		Pattern:     "/api/dogs/{id}",
		HandlerFunc: handler.GetDogByIDHandler, // TODO: add handler func
	},
	{
		Name:        "DeleteDogByID",
		Method:      http.MethodDelete,
		Pattern:     "/api/dogs/{id}",
		HandlerFunc: handler.DeleteDogByIDHandler, // TODO: add handler func
	},
	{
		Name:        "EditDogByID",
		Method:      http.MethodPut,
		Pattern:     "/api/dogs/{id}",
		HandlerFunc: handler.EditDogByIDHandler, // TODO: add handler func
	},
}
