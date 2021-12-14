/*
 * @Author: tangqimin
 * @Date: 2021-12-08 22:31:04
 * @Description:
 * @LastEditTime: 2021-12-14 22:18:43
 * @LastEditors: 汤启民
 * @FilePath: \todolist-golang\main.go
 */
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySQL() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = DB.DB().Ping()
	return
}

func main() {
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	// 程序结束后退出
	defer DB.Close()
	// 绑定模型
	DB.AutoMigrate(&Todo{})

	// init
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo", func(c *gin.Context) {
			var todo Todo
			c.BindJSON(&todo)
			if err = DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		v1Group.GET("/todo", func(c *gin.Context) {})
		v1Group.GET("/todo/:id", func(c *gin.Context) {})
		v1Group.PUT("/todo/:id", func(c *gin.Context) {})
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {})
	}
	r.Run(":9000")
}
