package main

import (
	"Project01/main/controller"
	"Project01/main/dao"
	"Project01/main/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	userDao        dao.UserDao               = dao.InitUserDao()
	userService    service.UserService       = service.New(userDao)
	userController controller.UserController = controller.New(userService)
)

func main() {
	defer userDao.CloseDB()
	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())
	apiGroup := server.Group("/user")
	{
		apiGroup.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"list":    userController.FindAll(),
				"message": "Success!"})
		})

		apiGroup.POST("/", func(context *gin.Context) {
			err := userController.Insert(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{"message": "Success!"})
			}
		})

		apiGroup.PUT("/:id", func(context *gin.Context) {
			err := userController.Update(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{"message": "Success!"})
			}
		})

		apiGroup.DELETE("/:id", func(context *gin.Context) {
			err := userController.Delete(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{"message": "Success!"})
			}
		})
	}
	server.Run()
}
