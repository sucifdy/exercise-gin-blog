package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var Posts = []Post{
	{ID: 1, Title: "Judul Postingan Pertama", Content: "Ini adalah postingan pertama di blog ini.", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 2, Title: "Judul Postingan Kedua", Content: "Ini adalah postingan kedua di blog ini.", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

func main() {
	router := SetupRouter()
	router.Run("localhost:8080")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/posts", getAllPosts)
	router.GET("/posts/:id", getPostByID)
	router.POST("/posts", createPost)

	return router
}

func getAllPosts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"posts": Posts})
}

func getPostByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}

	for _, post := range Posts {
		if post.ID == id {
			c.JSON(http.StatusOK, gin.H{"post": post})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Postingan tidak ditemukan"})
}

func createPost(c *gin.Context) {
	var newPost Post
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	newPost.ID = len(Posts) + 1
	newPost.CreatedAt = time.Now()
	newPost.UpdatedAt = time.Now()
	Posts = append(Posts, newPost)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Postingan berhasil ditambahkan",
		"post":    newPost,
	})
}
