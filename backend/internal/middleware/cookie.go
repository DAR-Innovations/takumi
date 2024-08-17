package middleware

import (
	"net/http"
	"takumi/internal/config"

	"github.com/gin-gonic/gin"
)

func SetCookieHandler(c *gin.Context, token string) {
	c.SetCookie("UserCookie", token, 3600, "/", config.HOST, true, true)
	c.String(http.StatusOK, "Cookie has been set")
}

func GetCookieHandler(c *gin.Context) {
	cookie, err := c.Cookie("UserCookie")
	if err != nil {
		c.String(http.StatusNotFound, "Cookie not found")
		return
	}
	c.String(http.StatusOK, "Cookie value: %s", cookie)
}

func DeleteCookieHandler(c *gin.Context) {
	c.SetCookie("UserCookie", "", -1, "/", config.HOST, true, true)
	c.String(http.StatusOK, "Cookie has been deleted")
}
