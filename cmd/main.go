package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/golang-rest-pagination/infrastructure"
	"github.com/prayogatriady/golang-rest-pagination/internal/app/handler"
	"github.com/prayogatriady/golang-rest-pagination/internal/app/repository"
	"github.com/prayogatriady/golang-rest-pagination/internal/app/service"
	"github.com/prayogatriady/golang-rest-pagination/internal/config"
)

func main() {

	cfg := config.NewConfig()
	
	db, err := infrastructure.NewDatabase(cfg)
	if err != nil {
		log.Fatal("Error initialize database:", err)
	}

	contactRepository := repository.NewUserRepository(db)
	contactService := service.NewContactService(contactRepository)
	contactHandler := handler.NewContactHandler(contactService)

	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/contacts", contactHandler.GetContactList)
		api.GET("/contact/:id", contactHandler.GetContact)
		api.POST("/contact", contactHandler.CreateContact)
		api.PUT("/contact/:id", contactHandler.UpdateContact)
		api.DELETE("/contact/:id", contactHandler.DeleteContact)
	}

	log.Fatal(router.Run(":" + cfg.App.Port))
}
