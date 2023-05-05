package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vaults-dev/vaults-backend/constants"
	"github.com/vaults-dev/vaults-backend/initializers"
	"github.com/vaults-dev/vaults-backend/models"
)

func ValidateAuth(c *gin.Context) {
	var response models.Response

	tokenString, err := c.Cookie("token")
	if err != nil {
		response.Error = "please login again"
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(constants.SECRETE_KEY), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp, _ := claims["exp"].(float64)

		if float64(time.Now().Unix()) > exp {
			response.Error = "token expired, please login again"
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		var user models.User
		initializers.DBconn.First(&user, "email=?", claims["sub"])
		if user.Email == "" {
			response.Error = "cookie data not valid, please login again"
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("user", user)

		c.Next()
	} else {
		response.Error = "please login again"
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

}
