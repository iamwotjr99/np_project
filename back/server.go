package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/iamwotjr99/np_project/back/api/middleware"
	"github.com/iamwotjr99/np_project/back/api/models"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Test
	r.GET("/someGet", someMethod)

	// User
	r.POST("/user/auth", authorizeUser)

	// Article
	// r.GET("/articles", articleAll)
	// r.POST("/articles", articleAdd)
	r.POST("/upload", uploadHandler)
	r.GET("/test", testGetAll)
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

func testGetAll(c *gin.Context) {
	fmt.Println("test start!")
	articles, err := models.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
	})

	fmt.Println("test end!")
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

// func articleAll(c *gin.Context) {
// 	a := models.Article{}
// 	articles, err := a.GetAll()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"result": articles,
// 		"count":  len(articles),
// 	})
// }

// func articleAdd(c *gin.Context) {
// 	var article models.Article

// 	if err := c.BindJSON(&article); err != nil {
// 		c.AbortWithStatus(http.StatusBadRequest)
// 	} else {
// 		models.Add(article)
// 		c.IndentedJSON(http.StatusCreated, gin.H{
// 			"msg": "Post OK",
// 		})
// 	}
// }

func uploadHandler(c *gin.Context) {
	var form models.Article

	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(&form)
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	fmt.Println(&form)

	// files := form.File["file"]
	images := []models.Image{}
	imageList := models.ImageList{Images: images}

	for _, file := range form.Images {
		fileExt := filepath.Ext(file.Filename)
		originalFileName := strings.TrimSuffix(filepath.Base(file.Filename), filepath.Ext(file.Filename))
		now := time.Now()
		filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
		filepath := "http://192.168.0.53:8008/image/" + filename

		image := models.Image{Filename: filename, Filepath: filepath}

		imageList.AddItem(image)

		out, err := os.Create("./images/" + filename)
		if err != nil {
			log.Fatal(err)
			c.String(http.StatusInternalServerError, "unknown error")
			return
		}
		defer out.Close()

		readerFile, _ := file.Open()
		_, err = io.Copy(out, readerFile)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(len(imageList.Images))

	models.Add(form, imageList)

	c.JSON(http.StatusOK, gin.H{"msg": "OK"})
}
