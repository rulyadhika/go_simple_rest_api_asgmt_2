package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/fga_digitalent_assignment_2/app"
	"github.com/rulyadhika/fga_digitalent_assignment_2/handler"
	"github.com/rulyadhika/fga_digitalent_assignment_2/helper"
	"github.com/rulyadhika/fga_digitalent_assignment_2/repository"
	"github.com/rulyadhika/fga_digitalent_assignment_2/service"
)

func main() {
	ginEngine := gin.Default()

	db := app.InitiateDB()
	orderRepository := repository.NewOrderRepositoryImpl()
	orderService := service.NewOrderServiceImpl(db, orderRepository)
	orderHandler := handler.NewOrderHandlerImpl(orderService)

	app.NewOrderRouter(ginEngine, orderHandler)

	srv := http.Server{
		Addr:    ":8080",
		Handler: ginEngine,
	}

	err := srv.ListenAndServe()

	helper.PanicIfErr(err)

}
