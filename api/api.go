package api

import (
	"encoding/json"
	"gitea.kood.tech/hannessoosaar/cars/pkg/models"
	"log"
	"net/http"
)

// Retrieves all the models data
func GetModels() []models.CarModel {
	response, err := http.Get("http://localhost:3000/api/models")
	if err != nil {
		log.Fatal("Error fetching data:", err)
	}
	defer response.Body.Close()
	var carModels []models.CarModel
	if err := json.NewDecoder(response.Body).Decode(&carModels); err != nil {
		log.Fatal("Error decoding response:", err)
	}
	return carModels
}

func GetManufacturer() []models.Manufacturer {
	response, err := http.Get("http://localhost:3000/api/manufacturers")
	if err != nil {
		log.Fatal("Error fetching data:", err)
	}
	defer response.Body.Close()
	var manufacturers []models.Manufacturer
	if err := json.NewDecoder(response.Body).Decode(&manufacturers); err != nil {
		log.Fatal("Error decoding response:", err)
	}
	return manufacturers
}

func GetCategory() []models.Category {
	response, err := http.Get("http://localhost:3000/api/categories")
	if err != nil {
		log.Fatal("Error fetching data:", err)
	}
	defer response.Body.Close()
	var categories []models.Category
	if err := json.NewDecoder(response.Body).Decode(&categories); err != nil {
		log.Fatal("Error decoding response:", err)
	}
	return categories
}

//TODO:Get models by id
//TODO:Get category by id
//TODO:Get Manufacturer by id
