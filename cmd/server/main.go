package main

import (
	"fmt"
	"net/http"

	"gitea.kood.tech/hannessoosaar/cars/intenal/conf"
	"gitea.kood.tech/hannessoosaar/cars/pkg/handle"
)

func main() {
	fs := http.FileServer(http.Dir("../../static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handle.LoadIndex)
	http.HandleFunc("/modelscompare", handle.LoadCompareModels)
	http.HandleFunc("/models", handle.LoadModels)
	http.HandleFunc("/model", handle.LoadModel)
	http.HandleFunc("/filter", handle.LoadFilter)
	fmt.Printf("Running on Port  :%s\n", conf.PORT)
	err := http.ListenAndServe(conf.PORT, nil)
	if err != nil {
		fmt.Printf("Error:%s", err)
	}
	fmt.Printf("Server started on Port: %s \n", conf.PORT)
}
