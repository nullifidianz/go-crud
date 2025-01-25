package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nullifidianz/go-crud/src/config/catch_err"
	"github.com/nullifidianz/go-crud/src/controller/models/request"
)

func CreateUser(c *gin.Context) {
	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		MyError := catch_err.NewBadRequestError(
			fmt.Sprintf("Some of the fields are missing or invalid. Please check and try again, error: %s", err.Error()))
		c.JSON(MyError.Code, MyError)
		return
	}
	fmt.Println(UserRequest)
}
