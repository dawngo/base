package baseController

import (
	"net/http"

	"github.com/Brave-man/base/bootstrap/database"
	"github.com/gin-gonic/gin"
)

// BaseController 基础Controller
type BaseController struct {
	DBSql database.DBSql
}

func GetIndex(c *gin.Context) {
	c.String(http.StatusOK, "Hello,World!")
}

func PostIndex(c *gin.Context) {
	// Login 登录信息
	type Login struct {
		User     string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}

	var json Login
	if err := c.BindJSON(&json); err == nil {
		if json.User == "david" && json.Password == "123" {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Illegal parameter"})
	}
}
