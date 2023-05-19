package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vaults-dev/vaults-backend/controllers"
	"github.com/vaults-dev/vaults-backend/initializers"
	"github.com/vaults-dev/vaults-backend/middlewares"
)

func init() {
	initializers.ConnectDB()

}

func main() {
	r := gin.Default()

	// r.Use(cors.New(
	// 	cors.Config{
	// 		AllowAllOrigins:  true,
	// 		AllowCredentials: true,
	// 		AllowMethods:     []string{"POST", "GET", "PUT", "OPTIONS"},
	// 	}))

	r.POST("/sign-up", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/home", middlewares.ValidateAuth, controllers.Home)
	r.Run() // listen and serve on 0.0.0.0:8080
}
