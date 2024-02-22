package main

import (
	"fmt"
	"net/http"

	"gitea.kood.tech/hannessoosaar/cars/intenal/conf"
	"gitea.kood.tech/hannessoosaar/cars/pkg/handle"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/",handle.GetIndex)
	err := http.ListenAndServe(conf.PORT, nil)
	if err != nil {
		fmt.Printf("Error:%s", err)
	}
	fmt.Printf("Server started on Port: %s \n", conf.PORT)
}
