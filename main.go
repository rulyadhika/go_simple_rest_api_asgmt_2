package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/go_simple_rest_api_asgmt_2/app"
	"github.com/rulyadhika/go_simple_rest_api_asgmt_2/exception"
	"github.com/rulyadhika/go_simple_rest_api_asgmt_2/handler"
	"github.com/rulyadhika/go_simple_rest_api_asgmt_2/helper"
	"github.com/rulyadhika/go_simple_rest_api_asgmt_2/repository"
	"github.com/rulyadhika/go_simple_rest_api_asgmt_2/service"
)

func main() {
	appConfig := app.GetAppConfig()
	ginEngine := gin.Default()
	ginEngine.Use(exception.ErrorHandler())

	db := app.InitiateDB()
	orderRepository := repository.NewOrderRepositoryImpl()
	orderService := service.NewOrderServiceImpl(db, orderRepository)
	orderHandler := handler.NewOrderHandlerImpl(orderService)

	app.NewOrderRouter(ginEngine, orderHandler)

	// run the server
	err := ginEngine.Run(":" + appConfig.SERVER_PORT)
	helper.PanicIfErr(err)
}
