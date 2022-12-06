package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
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
	r.GET("/article/:id", artcileGet)
	r.GET("/articles", articleGetAll)
	r.POST("/upload", addArticle)
	r.PUT("/article/:id", updateArticle)
	r.DELETE("/article/:id", deleteArticle)
	return r
}

func main() {
	r := setupRouter()
	r.Use(middleware.CORS)
	r.Static("/image", "./images")
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

func artcileGet(c *gin.Context) {
	strId := c.Params.ByName("id")

	intId, err := strconv.Atoi(strId)
	if err != nil {
		panic(err.Error())
	}

	article, err := models.Get(intId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"article": article,
	})
}

func articleGetAll(c *gin.Context) {
	articles, err := models.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
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

func addArticle(c *gin.Context) {
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
		filePath := "http://192.168.0.53:8008/image/" + filename
		// filePath := "http://192.168.25.52:8008/image/" + filename

		image := models.Image{Filename: filename, Filepath: filePath}

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

	c.JSON(http.StatusOK, gin.H{"msg": "Post OK"})
}

func updateArticle(c *gin.Context) {
	strId := c.Params.ByName("id")

	intId, err := strconv.Atoi(strId)
	if err != nil {
		panic(err.Error())
	}

	var form models.Article

	err = c.ShouldBind(&form)
	if err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	images := []models.Image{}
	imageList := models.ImageList{Images: images}

	for _, file := range form.Images {
		fileExt := filepath.Ext(file.Filename)
		originalFileName := strings.TrimSuffix(filepath.Base(file.Filename), filepath.Ext(file.Filename))
		now := time.Now()
		fileName := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
		filePath := "http://192.168.0.53:8008/image/" + fileName
		// filePath := "http://192.168.25.52:8008/image/" + filename

		image := models.Image{Filename: fileName, Filepath: filePath}

		imageList.AddItem(image)

		out, err := os.Create("./images/" + fileName)
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

	models.Update(form, imageList, intId)

	c.JSON(http.StatusOK, gin.H{"msg": "Update OK"})

}

func deleteArticle(c *gin.Context) {
	strId := c.Params.ByName("id")
	fmt.Println(strId)

	intId, err := strconv.Atoi(strId)
	if err != nil {
		panic(err.Error())
	}

	err = models.Delete(intId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Delete OK",
	})
}
