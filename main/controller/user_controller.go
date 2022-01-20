package controller

import (
	"Project01/main/models"
	"Project01/main/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
)

type UserController interface {
	Insert(c *gin.Context) error
	Update(c *gin.Context) error
	Delete(c *gin.Context) error
	FindAll() []models.User
	ShowAll(c *gin.Context)
}

type userController struct {
	userService service.UserService
}

var validate *validator.Validate

func (controller *userController) Insert(c *gin.Context) error {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		return err
	}
	err = validate.Struct(user)
	if err != nil {
		return err
	}
	controller.userService.Insert(user)
	return nil
}

func (controller *userController) Update(c *gin.Context) error {
	var user models.User
	id, err := strconv.ParseInt(c.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	user.Id = int(id)
	controller.userService.Update(user)
	return nil
}

func (controller *userController) Delete(c *gin.Context) error {
	var user models.User
	id, err := strconv.ParseInt(c.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	user.Id = int(id)
	controller.userService.Delete(user)
	return nil
}

func (controller *userController) FindAll() []models.User {
	return controller.userService.FindAll()
}

func (controller *userController) ShowAll(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func New(userService service.UserService) UserController {
	validate = validator.New()
	return &userController{
		userService: userService,
	}
}
