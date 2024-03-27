package handle

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"gitea.kood.tech/hannessoosaar/cars/api"
	"gitea.kood.tech/hannessoosaar/cars/pkg/models"
)

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

func LoadModel(w http.ResponseWriter, r *http.Request){
	modeID := r.URL.Query().Get("modelID")
	carModel := api.GetModelBy(modeID)
	carManufacturer :=api.GetManufacturerBy(strconv.Itoa(carModel.ManufacturerID))
	carCategory := api.GetCategoryBy(strconv.Itoa(carModel.CategoryId))
	fmt.Println(modeID)

	data :=struct {
		CarModel models.CarModel
		CarManufacturer models.Manufacturer
		CarCategory models.Category
	}{
		CarModel: carModel,
		CarManufacturer: carManufacturer,
		CarCategory: carCategory,
	}
	pageHtml,err := template.ParseFiles("../../template/model.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error getting template:", err)
		return
	}
	err = pageHtml.Execute(w,data)
	if err !=nil  {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func LoadFilter(w http.ResponseWriter, r *http.Request){
	pageHtml, err := template.ParseFiles("../../template/index.html")
	if err != nil {
		http.Error(w, err(), http.StatusInternalServerError)
		log.Println("Error getting template:", err)
		return
	}	
	if r.Method == http.MethodPost{
		manufactureId := r.FormValue("selectedManufacturerId")
		categoryId := r.FormValue("selectedCategoryId")
	cars := api.GetModelFilterBy(manufactureId,categoryId )
	}

	//TODO: Continue from here!
	
}
