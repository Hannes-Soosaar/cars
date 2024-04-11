package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"gitea.kood.tech/hannessoosaar/cars/pkg/models"
)

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

func GetModelBy(Id string) models.CarModel {
	response, err := http.Get("http://localhost:3000/api/models/" + Id)
	if err != nil {
		log.Fatal("Error fetching data:", err)
	}
	defer response.Body.Close()
	var car models.CarModel
	if err := json.NewDecoder(response.Body).Decode(&car); err != nil {
		log.Fatal("Error decoding response:", err)
	}
	return car
}

func FilterModelByCategory(categoryId string) (cars []models.CarModel) {
	allCars := GetModels()
	if len(categoryId) > 0 {
		for _, car := range allCars {
			if strconv.Itoa(car.CategoryId) == categoryId {
				cars = append(cars, car)
			}
		}
	} else {
		return allCars
	}
	return cars
}

func FilterModelByManufacturer(manufactureId string) (cars []models.CarModel) {
	allCars := GetModels()
	if len(manufactureId) > 0 {
		for _, car := range allCars {
			if strconv.Itoa(car.ManufacturerID) == manufactureId {
				cars = append(cars, car)
			}
		}
	} else {
		return allCars
	}
	return cars
}

func FilterModelByTransmission(transmission string) (cars []models.CarModel) {
	allCars := GetModels()
	if transmission == "" {
		return allCars
	}
	for _, car := range allCars {
		if car.Specs.Transmission == transmission {
			cars = append(cars, car)
		}
	}
	return cars
}

func GetManufacturerBy(Id string) models.Manufacturer {
	response, err := http.Get("http://localhost:3000/api/manufacturers/" + Id)
	if err != nil {
		log.Fatal("Error fetching data:", err)
	}
	defer response.Body.Close()
	var manufacturer models.Manufacturer
	if err := json.NewDecoder(response.Body).Decode(&manufacturer); err != nil {
		log.Fatal("Error decoding response:", err)
	}
	return manufacturer
}
func GetCategoryBy(Id string) models.Category {
	response, err := http.Get("http://localhost:3000/api/categories/" + Id)
	if err != nil {
		log.Fatal("Error fetching data:", err)
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


