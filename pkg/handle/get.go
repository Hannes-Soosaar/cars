package handle

import (
	"encoding/json"
	"fmt"
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
	err = pageHtml.Execute(w, pageHtml)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error executing template:", err)
		return
	}
	go GetModels()
	go GetManufacturer()
	go GetCategory()
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

func LoadCompareModels(w http.ResponseWriter, r *http.Request) {

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
	for _, car := range carModels {
		fmt.Println(car)
	}
	//TODO return a error if the car is not found
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
	for _, manufacturer := range manufacturers {
		fmt.Println(manufacturer)
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
	for _, category := range categories {
		fmt.Println(category)
	}
	return categories
}
