package users

import (
	"github.com/gin-gonic/gin"
	"github.com/laur1s/bookstore_users-api/domain/users"
	"net/http"
)

func CreateUser(c *gin.Context)  {
	var user users.User
	c.String(http.StatusNotImplemented, "Implement me!")
}

func GetUser(c *gin.Context)  {
	c.String(http.StatusNotImplemented, "Implement me!")
	
}

func SearchUser(c *gin.Context)  {
	c.String(http.StatusNotImplemented, "Implement me!")
}