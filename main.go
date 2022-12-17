package main

import (
	"github.com/gin-gonic/gin"
	"go_gin_todo/config"
	"go_gin_todo/models"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func setupRouter() *gin.Engine {

	//ログ設定
	f, _ := os.Create(config.Config.LogFile)
	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)

	//templateディレクトリ設定
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")

	// todo create
	r.POST("/todos/save", func(c *gin.Context) {

		models.CreateTodo(c.PostForm("content"))

		c.Redirect(http.StatusMovedPermanently, "/todos/list")
	})

	// todo create
	r.POST("/todos/update", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.PostForm("id"))
		content := c.PostForm("content")
		todo, _ := models.GetTodo(id)
		todo.Content = content
		models.UpdateTodo(todo)

		c.Redirect(http.StatusMovedPermanently, "/todos/list")
	})

	//todo edit
	r.GET("/todos/edit", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			log.Fatalln(err)
		}
		todo, _ := models.GetTodo(id)

		c.HTML(http.StatusOK, "edit.tmpl", gin.H{
			"title": "Todo",
			"todo":  todo,
		})
	})

	//todo delete
	r.GET("/todos/destroy", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			log.Fatalln(err)
		}
		models.DeleteTodo(id)

		c.Redirect(http.StatusMovedPermanently, "/todos/list")
	})

	r.GET("/todos/list", func(c *gin.Context) {

		var todos []models.Todo
		models.Db.Find(&todos)

		c.HTML(http.StatusOK, "list.tmpl", gin.H{
			"title": "Todo",
			"todos": todos,
		})
	})

	return r
}

func main() {

	r := setupRouter()
	r.Run(":8080")
}
