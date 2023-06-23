package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vaults-dev/vaults-backend/utils"
)

type JwkResponse struct {
	Keys []JwkResponseKey `json:"keys"`
}

type JwkResponseKey struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Use string `json:"use"`
	Alg string `json:"alg"`
	N   string `json:"n"`
	E   string `json:"e"`
}

func GetJwk(c *gin.Context) {
	jwk, err := utils.GetJwk()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp := JwkResponse{
		Keys: []JwkResponseKey{
			{
				Kty: "RSA",
				Kid: string(jwk.Kid),
				Use: "sig",
				Alg: "RS256",
				N:   string(jwk.N),
				E:   string(jwk.E),
			},
		},
	}

	c.Header("Content-Type", "application/json")
	c.JSON(
		http.StatusOK,
		resp,
	)
}
