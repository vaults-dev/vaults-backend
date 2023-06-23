package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vaults-dev/vaults-backend/initializers"
	"github.com/vaults-dev/vaults-backend/models"
	"github.com/vaults-dev/vaults-backend/utils"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	request := models.SignUp{}
	response := models.Response{}

	err := c.Bind(&request)
	if err != nil {
		response.Error = fmt.Sprintf("failed bind body, %v", err.Error())
		c.JSON(http.StatusBadRequest, response)

		return
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		response.Error = fmt.Sprintf("failed generate hash, %v", err.Error())
		c.JSON(http.StatusInternalServerError, response)

		return
	}

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(hashPass),
	}

	result := initializers.DBconn.Create(&user)
	if result.Error != nil {
		response.Error = fmt.Sprintf("failed create user to db, %v", result.Error.Error())
		c.JSON(http.StatusInternalServerError, response)

		return
	}

	response.Message = "success register"

	c.JSON(http.StatusCreated, response)
}

func Login(c *gin.Context) {

	// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	// c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	// c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	request := models.Login{}
	response := models.Response{}

	err := c.Bind(&request)
	if err != nil {
		response.Error = fmt.Sprintf("failed bind body, %v", err.Error())
		c.JSON(http.StatusBadRequest, response)

		return
	}

	var user models.User

	initializers.DBconn.First(&user, "email=?", request.Email)
	if user.Email == "" {
		response.Error = fmt.Sprintf("wrong email or pass")
		c.JSON(http.StatusBadRequest, response)

		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		response.Error = fmt.Sprintf("wrong email or pass")
		c.JSON(http.StatusBadRequest, response)

		return
	}

	jwt, err := utils.GenerateTokenForUser(user.Email)
	if err != nil {
		response.Error = fmt.Sprintf("failed generate jwt, %v", err.Error())
		c.JSON(http.StatusInternalServerError, response)

		return
	}

	response.Message = "success login"
	response.Data = struct {
		Jwt string `json:"jwt"`
	}{
		Jwt: string(jwt),
	}

	c.JSON(http.StatusOK, response)
}

func Home(c *gin.Context) {
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
