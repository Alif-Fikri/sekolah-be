package main

import (
	"sekolah-be/database"
	"sekolah-be/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Konek()

	r := gin.Default()
	routes.Api(r)

	r.Run()
}
