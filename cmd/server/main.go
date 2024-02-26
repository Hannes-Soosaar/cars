package main

import (
	"fmt"
	"net/http"
	"os"

	"gitea.kood.tech/hannessoosaar/cars/intenal/conf"
	"gitea.kood.tech/hannessoosaar/cars/pkg/handle"
)

func main() {

	http.Handle("/../../static/", http.StripPrefix("/../../static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handle.GetIndex)
	http.HandleFunc("/modelsCompare", handle.GetCompareModels)
	http.HandleFunc("/models", handle.GetModels)
	dir, _ := os.Getwd()
	fmt.Printf("working directory :%s\n", dir)
	fmt.Printf("Running on Port  :%s\n", conf.PORT)
	err := http.ListenAndServe(conf.PORT, nil)
	if err != nil {
		fmt.Printf("Error:%s", err)
	}
	fmt.Printf("Server started on Port: %s \n", conf.PORT)
}
