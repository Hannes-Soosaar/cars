package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gitea.kood.tech/hannessoosaar/cars/pkg/models"
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



func GetModelBy(Id string) models.CarModel{
	fmt.Println("http://localhost:3000/api/models/"+Id)
	response, err := http.Get("http://localhost:3000/api/models/"+Id)
	if 	err != nil {
		log.Fatal("Error fetching data:", err )
	}
	defer response.Body.Close()
	var car models.CarModel 
	if err := json.NewDecoder(response.Body).Decode(&car); err != nil {
		log.Fatal("Error decoding response:", err)
	}
	return car
}

func GetModelFilterBy(manufactureID string, categoryId string) (cars []models.CarModel) {

	return cars
}

func GetManufacturerBy(Id string) models.Manufacturer{
	fmt.Println("http://localhost:3000/api/manufacturers/"+Id)
	response, err := http.Get("http://localhost:3000/api/manufacturers/"+Id)
	if 	err != nil {
		log.Fatal("Error fetching data:", err )
	}
	defer response.Body.Close()
	var manufacturer models.Manufacturer 
	if err := json.NewDecoder(response.Body).Decode(&manufacturer); err != nil {
		log.Fatal("Error decoding response:", err)
	}
	return manufacturer
}
func GetCategoryBy(Id string) models.Category{
	fmt.Println("http://localhost:3000/api/categories/"+Id)
	response, err := http.Get("http://localhost:3000/api/categories/"+Id)
	if 	err != nil {
		log.Fatal("Error fetching data:", err )
	}
	defer response.Body.Close()
	var category models.Category 
	if err := json.NewDecoder(response.Body).Decode(&category); err != nil {
		log.Fatal("Error decoding response:", err)
	}
	return category
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