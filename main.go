package main

import (
	"github.com/gin-gonic/gin"
	"github.com/chawisali/gofinal/task"
	"github.com/chawisali/gofinal/middleware"
)


func setupRouter() *gin.Engine {
	r := gin.Default() 
	r.Use(middleware.Auth) 

	r.POST("/customers",task.CreateDetailHandler) 
	r.GET("/customers/:id",task.GetCustomerByIdHandler) 
	r.GET("/customers",task.GetAllCustomersHandler) 
	r.PUT("/customers/:id",task.UpdateDetailCustomerHandler) 
	r.DELETE("/customers/:id",task.DelCustomerHandler)  
	return r

}


func main() {
	//task.CreateTable() //First Step
	r := setupRouter()
	//run port ":2009"
	r.Run(":2009")
}