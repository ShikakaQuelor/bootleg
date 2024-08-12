package systembolaget

import (
	"context"
	"log"
	"shikakaQuelor/bootleg/internals/taxes"
	"sync"

	"github.com/alexgustafsson/systembolaget-api/v3/systembolaget"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type Metadata struct {
	DocumentCount int    `json:"docCount"`
	NextPage      int    `json:"nextPage"`
	Query         string `json:"query"`
	Total         int    `json:"total"`
}

type SearchResponse struct {
	Products []taxes.Product `json:"products"`
	MetaData Metadata        `json:"metaData"`
}

var (
	clientInstance *systembolaget.Client
	clientMutex    sync.Mutex
)

func Search(query string, page int, current int) (SearchResponse, gin.H) {
	client := getClient()

	searchOptions := &systembolaget.SearchOptions{
		Page:     page,
		PageSize: 10,
	}

	res, err := client.Search(
		context.Background(),
		searchOptions,
		systembolaget.FilterByQuery(query))

	if err != nil {
		client = swapClient()
		res, err = client.Search(
			context.Background(),
			searchOptions,
			systembolaget.FilterByQuery(query))
		if err != nil {
			return SearchResponse{}, gin.H{"error": err.Error()}
		}
	}

	var p []taxes.Product

	mapstructure.Decode(res.Products, &p)

	for k := range p {
		taxes.CalculateTaxesAndCut(&p[k])
	}

	log.Printf("Metainfo: %v", res.Metadata)

	return SearchResponse{
		Products: p,
		MetaData: Metadata{
			DocumentCount: res.Metadata.DocumentCount,
			NextPage:      res.Metadata.NextPage,
			Query:         query,
			Total:         current + len(p),
		}}, nil
}

func getClient() *systembolaget.Client {
	clientMutex.Lock()
	defer clientMutex.Unlock()

	if clientInstance == nil {
		clientInstance = createNewClient()
	}
	return clientInstance
}

func swapClient() *systembolaget.Client {
	clientMutex.Lock()
	defer clientMutex.Unlock()

	clientInstance = createNewClient()
	return clientInstance
}

func createNewClient() *systembolaget.Client {
	apiKey, err := systembolaget.GetAPIKey(context.Background())
	if err != nil {
		log.Fatalf("Failed to get API key: %v", err)
	}
	return systembolaget.NewClient(apiKey)
}
