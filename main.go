package main

import (
	db "github.com/dangyanglim/go-restful-api-simple/database"
	"github.com/gin-contrib/cors"
)

func main() {
	defer db.SqlDB.Close()
	router := initRouter()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"PUT", "PATCH", "GET", "POST", "DELETE"}
	router.Use(cors.New(config))
	router.Run(":8080")
}
