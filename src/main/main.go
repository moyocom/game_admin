package main

import (
	"app"
	_ "app/adminGame"
	_ "app/adminGamePlayer"
	_ "app/adminHome"
	_ "app/adminServer"
	_ "app/adminStats"
	_ "app/adminUser"
	"fmt"
	cfg "lib/config"
	"model"
	"net/http"
)

func main() {
	cfg.Load()

	fmt.Println("Server Start")
	model.Init()

	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.Handle("/", app.R)

	http.ListenAndServe(":8081", nil)

}
