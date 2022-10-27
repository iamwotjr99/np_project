package main

import (
	"fmt"
	"net/http"

	person "github.com/iamwotjr99/np_project/back/dto"

	"github.com/gin-gonic/gin"
)

func main() {
	p1 := person.Constructor("aa", "aa")
	fmt.Println(p1)
	name, pw := person.GetPerson(p1)
	fmt.Print(name, pw)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"pw":   pw,
		})
	})

	r.Run(":8008")

}
