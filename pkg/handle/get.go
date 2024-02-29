package handle

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"gitea.kood.tech/hannessoosaar/cars/pkg/models"
)

// TODO: Display all cars available
// TODO: Have a selecting table that sorts by Type
// TODO: Have a selecting table that sorts by Make
// TODO: Be able to save a ModelId in local storage
// TODO:

func LoadIndex(w http.ResponseWriter, r *http.Request) {
	pageHtml, err := template.ParseFiles("../../template/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error getting template:", err)
		return
	}
	availableModels := GetModels()
	availableManufacturers := GetManufacturer()
	availableCategories := GetCategory()
	data := struct {
		AvailableModels        []models.CarModel
		AvailableManufacturers []models.Manufacturer
		AvailableCategories    []models.Category
	}{
		AvailableModels:        availableModels,
		AvailableManufacturers: availableManufacturers,
		AvailableCategories:    availableCategories,
	}
	err = pageHtml.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LoadModels(w http.ResponseWriter, r *http.Request) {
	pageHtml, err := template.ParseFiles("../../template/models.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = pageHtml.Execute(w, pageHtml)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LoadManufacturer(w http.ResponseWriter, r *http.Request) {
	manufacturers := GetManufacturer()
	data := struct {
		ManufacturersInfo []models.Manufacturer
	}{
		ManufacturersInfo: manufacturers,
	}
	pageHtml, err := template.ParseFiles("../../template/manufacturer.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = pageHtml.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LoadCompareModels(w http.ResponseWriter, r *http.Request) {
	pageHtml, err := template.ParseFiles("../../template/modelscompare.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = pageHtml.Execute(w, pageHtml)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

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
