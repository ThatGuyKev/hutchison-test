package handler

import (
	"encoding/json"
	"hutchison-test/common"
	"hutchison-test/infrastructure"
	"hutchison-test/model"
	"hutchison-test/repository"
	"hutchison-test/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateDogHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Creating new dog...")

	newDog := &model.Dog{}

	if err := json.NewDecoder(r.Body).Decode(newDog); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if newDog.Variants != nil {

		variants, err := utils.CSVToJSONArray(*newDog.Variants)
		if err != nil {
			// TODO: handle edge case
		}

		newDog.Variants = &variants
	}

	dogsRepository := repository.DogsRepository{
		Db: infrastructure.DBAdapter,
	}

	createdDog, err := dogsRepository.Create(newDog)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response := &model.GenericResponse[model.CreateDogResponseData]{
		Data: model.CreateDogResponseData{
			Result:  true,
			Message: "Dog created successfully",
			Dog:     createdDog,
		},
		Meta: "Meta data",
	}
	common.RespondWithJSON(w, http.StatusOK, response)
}

func ListDogsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Listing Dogs...")
	dogsRepository := repository.DogsRepository{
		Db: infrastructure.DBAdapter,
	}
	dogsList, err := dogsRepository.ListAll()
	if err != nil {
		log.Printf("Error: %v ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response := &model.GenericResponse[model.ListDogsResponseData]{
		Data: model.ListDogsResponseData{
			Result:  true,
			Message: "Dogs List",
			Dogs:    dogsList,
		},
		Meta:   "Meta data",
		Error_: "",
	}
	common.RespondWithJSON(w, http.StatusOK, response)
}

func GetDogByIDHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Finding Dog By ID...")
	dogsRepository := repository.DogsRepository{
		Db: infrastructure.DBAdapter,
	}

	vars := mux.Vars(r)
	dogIdStr, ok := vars["id"]
	if !ok || dogIdStr == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	dogId, err := strconv.ParseUint(dogIdStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	foundDog, err := dogsRepository.GetByID(uint(dogId))
	if err != nil {
		log.Printf("Error: %v ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response := &model.GenericResponse[model.GetDogByIDResponseData]{
		Data: model.GetDogByIDResponseData{
			Result:  true,
			Message: "Found Dog",
			Dog:     foundDog,
		},
		Meta:   "Meta data",
		Error_: "",
	}
	common.RespondWithJSON(w, http.StatusOK, response)
}

func DeleteDogByIDHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Deleting Dog By ID...")
	dogsRepository := repository.DogsRepository{
		Db: infrastructure.DBAdapter,
	}

	vars := mux.Vars(r)
	dogIdStr, ok := vars["id"]
	if !ok || dogIdStr == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	dogId, err := strconv.ParseUint(dogIdStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = dogsRepository.DeleteByID(uint(dogId))
	if err != nil {
		log.Printf("Error: %v ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response := &model.GenericResponse[model.DeleteDogByIDResponseData]{
		Data: model.DeleteDogByIDResponseData{
			Result:  true,
			Message: "Dog deleted successfully",
		},
		Meta:   "Meta data",
		Error_: "",
	}
	common.RespondWithJSON(w, http.StatusOK, response)
}

func EditDogByIDHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Editing Dog By ID...")

	updatedDog := &model.Dog{}
	if err := json.NewDecoder(r.Body).Decode(updatedDog); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if updatedDog.Variants != nil {
		variants, err := utils.CSVToJSONArray(*updatedDog.Variants)
		if err != nil {
			// TODO: handle edge case
			updatedDog.Variants = nil
		} else {
			updatedDog.Variants = &variants
		}
	}

	dogsRepository := repository.DogsRepository{
		Db: infrastructure.DBAdapter,
	}

	vars := mux.Vars(r)
	dogIdStr, ok := vars["id"]
	if !ok || dogIdStr == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	dogId, err := strconv.ParseUint(dogIdStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = dogsRepository.EditByID(uint(dogId), updatedDog)
	if err != nil {
		log.Printf("Error: %v ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response := &model.GenericResponse[model.EditDogByIDResponseData]{
		Data: model.EditDogByIDResponseData{
			Result:  true,
			Message: "Dog edited successfully",
			Dog:     updatedDog,
		},
		Meta:   "Meta data",
		Error_: "",
	}
	common.RespondWithJSON(w, http.StatusOK, response)
}
