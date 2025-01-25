package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nullifidianz/go-crud/src/controller"
)

func InitRoutes(route *gin.RouterGroup) {
	route.GET("/getUserById/:userID", controller.FindUserById)
	route.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	route.POST("/createUser", controller.CreateUser)
	route.PUT("/updateUser/:userID", controller.UpdateUser)
	route.DELETE("/deleteUser/:userID", controller.DeleteUser)
}
