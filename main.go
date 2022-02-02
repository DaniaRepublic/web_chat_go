package main

import (
	"github.com/DaniaRepublic/commonSpaceGo/dbconnector"
	//"github.com/DaniaRepublic/commonSpaceGo/ejabber"
	"github.com/DaniaRepublic/commonSpaceGo/endpoints"
	"github.com/DaniaRepublic/commonSpaceGo/jwt"

	"github.com/gin-gonic/gin"
)

func main() {
	// connect to cache and relational databases
	// store them in environment for endpoints
	rdb := new(dbconnector.RedisConn)
	rdb.Connect()
	db := new(dbconnector.MYSQLConn)
	db.Connect()
	endpointsENV := &endpoints.Env{
		SQLDB: db,
		RDB:   rdb,
	}

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/login", endpoints.GETlogin)
	router.POST("/login", endpointsENV.POSTlogin)

	// group that requires authentication
	auth := router.Group("/")
	auth.Use(jwt.VerifyToken())
	{
		auth.GET("/main", endpointsENV.GETmain)
		auth.POST("/send", endpoints.POSTsend)
	}

	router.Run(":8080")
}
