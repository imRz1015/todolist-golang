/*
 * @Author: tangqimin
 * @Date: 2021-12-08 22:31:04
 * @Description:
 * @LastEditTime: 2021-12-08 23:01:19
 * @LastEditors: tangqimin
 * @FilePath: \studyGin\main.go
 */
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func main() {
	// todo : create database

	// todo : connect database

	// init
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo", func(c *gin.Context) {})
		v1Group.GET("/todo", func(c *gin.Context) {})
		v1Group.GET("/todo/:id", func(c *gin.Context) {})
		v1Group.PUT("/todo/:id", func(c *gin.Context) {})
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {})
	}
	r.Run(":9000")
}
