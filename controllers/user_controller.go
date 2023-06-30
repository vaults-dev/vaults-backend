package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vaults-dev/vaults-backend/libraries"
	"github.com/vaults-dev/vaults-backend/models"
)

type UserController struct {
	lib libraries.UserLibraryInterface
}

type UserControllerInterface interface {
	SignUp(c *gin.Context)
	Login(c *gin.Context)
	Home(c *gin.Context)
}

func NewUserController(lib libraries.UserLibraryInterface) UserControllerInterface {
	return &UserController{
		lib: lib,
	}
}

func (ctrl *UserController) SignUp(c *gin.Context) {
	request := models.SignUp{}
	response := models.Response{}

	err := c.Bind(&request)
	if err != nil {
		response.Error = fmt.Sprintf("failed bind body, %v", err.Error())
		c.JSON(http.StatusBadRequest, response)

		return
	}

	response.Error = ctrl.lib.SignUp(request)
	response.Message = "success register"

	c.JSON(http.StatusCreated, response)
}

func (ctrl *UserController) Login(c *gin.Context) {
	request := models.Login{}
	response := models.Response{}

	err := c.Bind(&request)
	if err != nil {
		response.Error = fmt.Sprintf("failed bind body, %v", err.Error())
		c.JSON(http.StatusBadRequest, response)

		return
	}

	jwt, err := ctrl.lib.Login(request)
	if err != nil {
		response.Error = err
		c.JSON(http.StatusBadRequest, response)

		return
	}

	response.Message = "success login"
	response.Data = struct {
		Jwt string `json:"jwt"`
	}{
		Jwt: jwt.(string),
	}

	c.JSON(http.StatusOK, response)
}

func (ctrl *UserController) Home(c *gin.Context) {
	var response models.Response

	userData, exist := c.Get("user")
	if !exist {
		response.Message = "Please Login again"
		c.JSON(http.StatusUnauthorized, response)
	}

	user, _ := userData.(models.User)
	response.Message = fmt.Sprintf("WELCOME HOME, %v", user.Name)

	c.JSON(http.StatusOK, response)
}
