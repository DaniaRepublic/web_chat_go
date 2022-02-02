package endpoints

import (
	"database/sql"
	"net/http"

	"github.com/DaniaRepublic/commonSpaceGo/classes"
	"github.com/DaniaRepublic/commonSpaceGo/dbconnector"
	"github.com/DaniaRepublic/commonSpaceGo/jwt"

	"github.com/gin-gonic/gin"
)

type Env struct {
	SQLDB *dbconnector.MYSQLConn // pointer to mysql
	RDB   *dbconnector.RedisConn // pointer to redis
}

const JWT_TTL int = 3600 // 3600 seconds = 1 hour

// endpoint handler naming convention: <METHOD><endpoint>
func GETlogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func (e *Env) POSTlogin(c *gin.Context) {
	var user classes.User

	jid := c.PostForm("jid")
	phone := c.PostForm("phone")

	user, err := e.SQLDB.UserByJIDandPhone(jid, phone)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error sql.ErrNoRows": err.Error(),
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error in UserByJIDandPhone": err.Error(),
			})
		}
	} else {
		token, err := jwt.GenerateToken(user, JWT_TTL)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error generating token": err.Error(),
			})
		} else {
			// finaly, redirect to
			c.SetCookie("JWTAuth", token, JWT_TTL, "/", "localhost", false, true)
			c.Redirect(http.StatusFound, "/main")
		}
	}
}

func (e *Env) GETmain(c *gin.Context) {
	jid, ok1 := c.Get("jid")
	username, ok2 := c.Get("username")
	if ok1 && ok2 {
		c.HTML(http.StatusOK, "main.html", gin.H{
			"jid":      jid,
			"username": username,
		})
	} else {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"error": "username or jid not set.",
		})
	}
}

func POSTsend(c *gin.Context) {

}
