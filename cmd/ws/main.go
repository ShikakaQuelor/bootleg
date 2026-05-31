package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"shikakaQuelor/bootleg/internals/systembolaget"
	"shikakaQuelor/bootleg/views"
	"shikakaQuelor/bootleg/views/components"

	"github.com/a-h/templ"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.SetTrustedProxies([]string{"127.0.0.1", "::1"})

	port, found := os.LookupEnv("PORT")
	if !found {
		port = "8080"
	}

	router.POST("/search", getSearch)
	router.POST("/details", getDetails)
	router.GET("/", getIndex)
	router.Static("/static", "./static/")

	router.Run(":" + port)
}

func getSearch(c *gin.Context) {
	data, err := formData(c)

	if err != nil {
		println(fmt.Sprintf("Error: %v", err))
		return
	}
	c.Header("HX-Push-Url", fmt.Sprintf("?search=%s", c.PostForm(("search"))))

	render(c, http.StatusOK, components.Products(data.Products, data.MetaData))
}

func getDetails(c *gin.Context) {
	var p components.ProductDetails

	err := c.ShouldBind(&p)

	if err != nil {
		fmt.Println(err)
	}

	render(c, http.StatusOK, components.Details(p))
}

func getIndex(c *gin.Context) {
	data, err := getQueryResults(c)
	if err != nil {
		return
	}
	render(c, http.StatusOK, views.Index(data.Products, data.MetaData))
}

func formData(c *gin.Context) (systembolaget.SearchResponse, gin.H) {
	query, found := c.GetPostForm("search")
	currentStr := c.DefaultPostForm("current", "0")
	nextPageStr := c.DefaultPostForm("nextPage", "1")

	if !found {
		return systembolaget.SearchResponse{}, nil
	}

	nextPage, err := strconv.Atoi(nextPageStr)
	if err != nil {
		println("Shit went to hell")
		return systembolaget.SearchResponse{}, nil
	}

	current, err := strconv.Atoi(currentStr)
	if err != nil {
		println("Shit went to hell")
		return systembolaget.SearchResponse{}, nil
	}

	return systembolaget.Search(query, nextPage, current)
}

func getQueryResults(c *gin.Context) (systembolaget.SearchResponse, gin.H) {
	query := c.Query("search")
	pageStr := c.DefaultQuery("page", "1")

	if query == "" {
		return systembolaget.SearchResponse{}, nil
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return systembolaget.SearchResponse{}, gin.H{"error": "Invalid page parameter"}
	}

	return systembolaget.Search(query, page, 0)
}

func render(ctx *gin.Context, status int, template templ.Component) error {
	ctx.Status(status)
	return template.Render(ctx.Request.Context(), ctx.Writer)
}
