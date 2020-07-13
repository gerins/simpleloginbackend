package main

import (
	"login_page_gerin/config/database"
	"login_page_gerin/config/router"
	"login_page_gerin/middleware"
)

func main() {
	db := database.ConnectDB()
	r := router.CreateRouter()
	r.Use(middleware.AccessCORS)
	router.NewAppRouter(db, r).InitRouter()
	router.StartServer(r)
}
