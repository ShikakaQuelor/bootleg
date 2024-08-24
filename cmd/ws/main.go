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

type VariableResponse struct {
	Systembolaget Systembolaget `json:"systembolaget"`
	Skatteverket  Skatteverket  `json:"skatteverket"`
}

type Systembolaget struct {
	PercentageCut float32  `json:"percentageCut"`
	UnitCuts      UnitCuts `json:"unitcuts"`
}

type UnitCuts struct {
	Spirits           float32 `json:"spirits"`
	Wine              float32 `json:"wine"`
	Cider             float32 `json:"cider"`
	Beer              float32 `json:"beer"`
	WineNonAlcoholic  float32 `json:"wineNonAlcoholic"`
	CiderNonAlcoholic float32 `json:"ciderNonAlcoholic"`
}

type Skatteverket struct {
	Beer        []TaxInfo `json:"beer"`
	Wine        []TaxInfo `json:"wine"`
	MiddleClass []TaxInfo `json:"middleClass"`
	Spirits     []TaxInfo `json:"spirits"`
}

type TaxInfo struct {
	AlcFrom               float32 `json:"alcFrom"`
	AlcTo                 float32 `json:"alc"`
	TaxPerPercentAndLiter float32 `json:"taxPerPercentAndLiter"`
}

var Variables = VariableResponse{
	Systembolaget: Systembolaget{
		PercentageCut: 0.147,
		UnitCuts: UnitCuts{
			Spirits:           6.37,
			Wine:              5.40,
			Cider:             1.52,
			Beer:              0.95,
			WineNonAlcoholic:  5.92,
			CiderNonAlcoholic: 2.46,
		},
	},
	Skatteverket: Skatteverket{
		Beer: []TaxInfo{
			{AlcFrom: 0, AlcTo: 2.8, TaxPerPercentAndLiter: 0},
			{AlcFrom: 2.8, AlcTo: 100, TaxPerPercentAndLiter: 2.28},
		},
		Wine: []TaxInfo{
			{AlcFrom: 0, AlcTo: 2.25, TaxPerPercentAndLiter: 0},
			{AlcFrom: 2.25, AlcTo: 4.5, TaxPerPercentAndLiter: 10.38},
			{AlcFrom: 4.5, AlcTo: 7, TaxPerPercentAndLiter: 15.34},
			{AlcFrom: 7, AlcTo: 8.5, TaxPerPercentAndLiter: 21.12},
			{AlcFrom: 8.5, AlcTo: 15, TaxPerPercentAndLiter: 29.58},
			{AlcFrom: 15, AlcTo: 18, TaxPerPercentAndLiter: 61.90},
		},
		MiddleClass: []TaxInfo{
			{AlcFrom: 0, AlcTo: 15, TaxPerPercentAndLiter: 37.34},
			{AlcFrom: 15, AlcTo: 22, TaxPerPercentAndLiter: 61.90},
		},
		Spirits: []TaxInfo{
			{AlcFrom: 0, AlcTo: 1.2, TaxPerPercentAndLiter: 0},
			{AlcFrom: 1.2, AlcTo: 100, TaxPerPercentAndLiter: 526.97},
		},
	},
}

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
