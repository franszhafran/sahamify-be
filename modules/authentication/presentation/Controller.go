package presentation

import (
	"github.com/franszhafran/sahamify-be/modules/authentication/application/service/login"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Controller is collected here
type Controller struct{}

// Login action
func Login(c *gin.Context) {
	request := login.NewRequest(c.DefaultPostForm("username", ""), c.DefaultPostForm("email", ""))

	if resolver, err := resolve("LoginService"); err {
		returnError(c)
		return
	} else {
		service := resolver.(*login.Service)
		service.Execute(request)

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "login",
		})
	}
}

func resolve(i string) (interface{}, bool) {
	if i == "LoginService" {
		return &login.Service{}, false
	}

	return nil, true
}

func returnError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  "failed",
		"message": "login",
	})
}
