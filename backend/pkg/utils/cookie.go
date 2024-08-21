package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CookieSettings struct {
	Name     string
	Data     string
	MaxAge   int
	Path     string
	Host     string
	Secure   bool
	HttpOnly bool
}

func SetCookieHandler(c *gin.Context, s CookieSettings) {
	c.SetCookie(s.Name, s.Data, s.MaxAge, s.Path, s.Host, s.Secure, s.HttpOnly)
	c.String(http.StatusOK, "\nCookie has been set")
}

func GetCookieHandler(c *gin.Context, s CookieSettings) (*string, error) {
	cookie, err := c.Cookie(s.Name)
	if err != nil {
		c.String(http.StatusNotFound, "\nCookie not found")
		return nil, err
	}
	c.String(http.StatusOK, "\nCookie value: %s", cookie)
	return &cookie, nil
}

func DeleteCookieHandler(c *gin.Context, s CookieSettings) {
	c.SetCookie(s.Name, "", -1, s.Path, s.Host, s.Secure, s.HttpOnly)
	c.String(http.StatusOK, "\nCookie has been deleted")
}
