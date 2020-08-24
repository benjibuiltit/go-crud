package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Post ...
type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

// Posts ...
type Posts []Post

// Mem DB ...
var db = Posts{
	Post{ID: "1", Title: "My title", Content: "This is my first post", Author: "Benji"},
	Post{ID: "2", Title: "Second Post", Content: "This is my second post", Author: "Benji"},
	Post{ID: "3", Title: "Third Post", Content: "This is my third post", Author: "Benji"},
}

// e.GET("/posts", listPosts)
func listPosts(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, db, "	")
}

// e.POST("/posts", createPost)
func createPost(c echo.Context) (err error) {
	post := new(Post)
	if err = c.Bind(post); err != nil {
		return
	}

	db = append(db, *post)
	return c.JSONPretty(http.StatusOK, db, "	")
}

// e.GET("/posts/:id", getPost)
func getPost(c echo.Context) error {
	id := c.Param("id")

	for _, post := range db {
		if post.ID == id {
			return c.JSONPretty(http.StatusOK, post, "	")
		}
	}

	return c.String(http.StatusNotFound, "Not Found")
}

// e.DELETE("/posts/:id", deletePost)
func deletePost(c echo.Context) error {
	id := c.Param("id")

	for postIndex, post := range db {
		if post.ID == id {
			db = append(db[:postIndex], db[postIndex+1:]...)
		}
	}

	return c.String(http.StatusOK, fmt.Sprintf("Post %s succesfully deleted", id))
}

// e.PUT("/posts/:id", updatePost)
func updatePost(c echo.Context) (err error) {
	id := c.Param("id")

	newPost := new(Post)
	if err = c.Bind(newPost); err != nil {
		return c.String(http.StatusBadRequest, "Malformed Request Body")
	}

	if id != newPost.ID {
		return c.String(http.StatusBadRequest, "Path ID and request body do not match")
	}

	for postIndex, post := range db {
		if post.ID == newPost.ID {
			db[postIndex] = *newPost
			return c.JSONPretty(http.StatusOK, newPost, "	")
		}
	}

	return c.String(http.StatusNotFound, "Not Found")

}

func main() {
	// New echo instance
	e := echo.New()

	// Register middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	// Register routes
	e.GET("/posts", listPosts)
	e.POST("/posts", createPost)
	e.GET("/posts/:id", getPost)
	e.DELETE("/posts/:id", deletePost)
	e.PUT("/posts/:id", updatePost)
	e.Logger.Fatal(e.Start(":1323"))
}
