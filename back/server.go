package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/iamwotjr99/np_project/back/api/middleware"
	"github.com/iamwotjr99/np_project/back/api/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Test
	r.GET("/someGet", someMethod)

	// User
	r.POST("/user/auth", authorizeUser)

	// Article
	r.GET("/articles", articleAll)
	r.POST("/articles", articleAdd)
	return r
}

func main() {
	r := setupRouter()
	r.Use(middleware.CORS)
	r.Run(":8008")

}

func someMethod(c *gin.Context) {
	httpMethod := c.Request.Method
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	a := c.ClientIP()
	fmt.Println("Request IP: ", a)
	c.JSON(http.StatusOK, gin.H{
		"status":  "good",
		"sending": httpMethod,
		"client":  a,
	})
}

func authorizeUser(c *gin.Context) {
	// var person Person
	// httpMethod := c.Request.Method
	var data models.Person
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
	} else {
		if data.Name == "wotjr" && data.Pw == "1234" {
			c.JSON(http.StatusOK, gin.H{
				// "sending": httpMethod,
				"data": data,
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid user",
			})
		}
	}
}

func articleAll(c *gin.Context) {
	a := models.Article{}
	articles, err := a.GetAll()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": articles,
		"count":  len(articles),
	})
}

func articleAdd(c *gin.Context) {
	var article models.Article

	if err := c.BindJSON(&article); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.Add(article)
		c.IndentedJSON(http.StatusCreated, gin.H{
			"msg": "Post OK",
		})
	}
}
