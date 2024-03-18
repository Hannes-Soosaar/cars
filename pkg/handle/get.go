package handle

import (
	"gitea.kood.tech/hannessoosaar/cars/api"
	"gitea.kood.tech/hannessoosaar/cars/pkg/models"
	"html/template"
	"log"
	"net/http"
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
	availableModels := api.GetModels()
	availableManufacturers := api.GetManufacturer()
	availableCategories := api.GetCategory()
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

func GetManufacturer() {
	panic("unimplemented")
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
	manufacturers := api.GetManufacturer()
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
