package main

import (
	l "Effective-Mobile/internal/dbconn"
	"Effective-Mobile/pkg/logger"
	"log"
	"net/http"

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
//swag init -g cmd/main.go --parseDependency --parseInternal -d ./,internal/structures,pkg/handlers && go run cmd/main.go - to start

func init() {
	l.Db = l.Connection()
	log.Println("PostgreSQL DB connected")
	logger.LogFile = logger.LoggerInit()
	log.Println("Started logging to logger.txt")
}

func main() {
	defer l.Db.Close()
	defer logger.LogFile.Close()
	log.SetOutput(logger.LogFile)
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/info", h.GetCars)
	r.PUT("/cars/:id", h.UpdateCar)
	r.POST("/cars", h.CreateCar)
	r.POST("/owners", h.AddOwner)
	r.DELETE("/cars/:id", h.DeleteCar)

	r.GET("/docs", func(c *gin.Context) { c.Redirect(http.StatusFound, "swagger/index.html") })
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
		return
	}

}
