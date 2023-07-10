package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type t struct {
	username string
	password string
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /login [get]
func Login(c *gin.Context) {

	name := c.Query("username")
	password := c.Query("password")
	if name != "admin" || password != "pwd" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "用户名密码错误",
		})

		c.AbortWithError(http.StatusInternalServerError, errors.New("login failed"))
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// @BasePath /api/v2

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /register [post]
func Regist(c *gin.Context) {

	data := t{
		username: c.Query("username"),
		password: "pwd",
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
