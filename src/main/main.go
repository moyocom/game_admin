package main

import (
	"app"
	_ "app/adminGame"
	_ "app/adminGamePlayer"
	_ "app/adminHome"
	_ "app/adminUser"
	"fmt"
	"model"
	"net/http"
)

func main() {
	fmt.Println("Server Start")

	model.Init()
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.Handle("/", app.R)
	http.ListenAndServe(":8081", nil)
}
