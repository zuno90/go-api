package routes

import (
	"api/controllers"

	"fmt"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine)  {
    fmt.Println("router")
    // All routes related to users comes here
    router.GET("/users", controllers.GetUsers())
    router.GET("/user/:id", controllers.GetUserById())
    router.POST("/user", controllers.CreateUser())
}