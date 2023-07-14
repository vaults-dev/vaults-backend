package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vaults-dev/vaults-backend/models"
	"github.com/vaults-dev/vaults-backend/utils"

	"golang.org/x/oauth2"
)

var (
	GoogleOauthConfig *oauth2.Config
	// TODO: randomize it
	oauthStateString = "pseudo-random"
)

func LoginPage(c *gin.Context) {
	var htmlIndex = `<html>
<body>
	<a href="/google-oauth">Google Log In</a>
</body>
</html>`

	fmt.Fprintf(c.Writer, htmlIndex)
}

func GoogleOAuth(c *gin.Context) {
	log.Println("-----------------------------------")
	log.Printf("GOOGLE CONFIG %+v\n", GoogleOauthConfig)
	ClientID := os.Getenv("GOOGLE_CLIENT_ID")
	log.Println("CLIENT ID is:", ClientID)

	url := GoogleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(c.Writer, c.Request, url, http.StatusTemporaryRedirect)
}

func GoogleCallback(c *gin.Context) {
	googleResp, err := getUserInfo(c.Request.FormValue("state"), c.Request.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(c.Writer, c.Request, "/login-page", http.StatusTemporaryRedirect)
		return
	}

	jwt, err := utils.GenerateTokenForUser(googleResp.Email)
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(c.Writer, c.Request, "/login-page", http.StatusTemporaryRedirect)
		return
	}

	// response := models.Response{
	// 	Message: "success login",
	// 	Data: struct {
	// 		Jwt string `json:"jwt"`
	// 	}{
	// 		Jwt: string(jwt),
	// 	},
	// }

	// c.JSON(http.StatusOK, response)

	http.Redirect(c.Writer, c.Request, fmt.Sprintf("/?jwt=%v", string(jwt)), http.StatusTemporaryRedirect)

}

func getUserInfo(state string, code string) (models.GoogleOAuthResponse, error) {
	googleResp := models.GoogleOAuthResponse{}

	if state != oauthStateString {
		return googleResp, fmt.Errorf("invalid oauth state")
	}

	token, err := GoogleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return googleResp, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return googleResp, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return googleResp, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	err = json.Unmarshal(contents, &googleResp)
	if err != nil {
		return googleResp, fmt.Errorf("failed unmarshal response body: %s", err.Error())
	}

	return googleResp, nil
}
