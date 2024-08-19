package middleware

import (
	"net/http"
	"takumi/internal/config"

	"github.com/gin-gonic/gin"
)

const TOKEN_KEY = "TakumiToken"

func SetCookieHandler(c *gin.Context, token string) {
	cfg := config.GetConfig()
	c.SetCookie(TOKEN_KEY, token, 3600, "/", cfg.FrontendUrl, false, true)
	c.String(http.StatusOK, "Cookie has been set")
}

func GetCookieHandler(c *gin.Context) {
	cookie, err := c.Cookie(TOKEN_KEY)
	if err != nil {
		c.String(http.StatusNotFound, "Cookie not found")
		return
	}
	c.String(http.StatusOK, "Cookie value: %s", cookie)
}

func DeleteCookieHandler(c *gin.Context) {
	cfg := config.GetConfig()
	c.SetCookie(TOKEN_KEY, "", -1, "/", cfg.FrontendUrl, false, true)
	c.String(http.StatusOK, "Cookie has been deleted")
}
