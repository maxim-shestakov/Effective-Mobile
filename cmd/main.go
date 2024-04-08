package main

import (
	"Effective-Mobile/internal/db"
	"log"
	"net/http"
	"os"

	_ "Effective-Mobile/docs"

	h "Effective-Mobile/pkg/handlers"

	gin "github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title Effective-Mobile API
// @version 1.0
// @description This is a RESTful API for Effective-Mobile project
// @host localhost:8080
// @BasePath /

//go install github.com/swaggo/swag/cmd/swag@v1.8.12
//swag init -g cmd/main.go --parseDependency --parseInternal -d ./,internal/db,pkg/handlers - to generate docs

func main() {
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dataBase := os.Getenv("POSTGRES_DB")

	if len(user) == 0 || len(pass) == 0 || len(host) == 0 || len(port) == 0 || len(dataBase) == 0 {
		log.Fatal("DB environment variables are not set")
	}
	db := db.New(user, pass, host, port, dataBase)
	log.Println("Connected to DB")
	defer db.Close()
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/info", h.GetCars(db))
	r.PUT("/cars/:id", h.UpdateCar(db))
	r.POST("/cars", h.CreateCar(db))
	r.POST("/owners", h.AddOwner(db))
	r.DELETE("/cars/:id", h.DeleteCar(db))

	r.GET("/docs", func(c *gin.Context) { c.Redirect(http.StatusFound, "swagger/index.html") })
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Server started")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
		return
	}

}
