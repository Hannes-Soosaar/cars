package handle

import (
	"html/template"
	"log"
	"net/http"
)

// TODO: Display all cars available
// TODO: Have a selecting table that sorts by Type
// TODO: Have a selecting table that sorts by Make
// TODO: Be able to save a ModelId in local storage
// TODO:
func GetIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("starting Get Index")
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
}

func GetModels(w http.ResponseWriter, r *http.Request) {
	pageHtml, err := template.ParseFiles("../../template/models.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = pageHtml.Execute(w, pageHtml)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetCompareModels(w http.ResponseWriter, r *http.Request) {

}
