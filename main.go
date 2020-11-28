package main

import (
	"github.com/asset-kingdom/assetkingdom/controller"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", controller.LoginHandler)
}
