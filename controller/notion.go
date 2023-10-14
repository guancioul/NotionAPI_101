package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guancioul/NotionGoogleCalendarIntegration/handler"
	"github.com/guancioul/NotionGoogleCalendarIntegration/infra"
	"github.com/guancioul/NotionGoogleCalendarIntegration/model/requestModel"
	"github.com/guancioul/NotionGoogleCalendarIntegration/model/responseModel"
	"github.com/guancioul/NotionGoogleCalendarIntegration/util"
)

// CreateNotionDatabase godoc
//
//	@Summary		Create a new Notion Database
//	@Description	Creates a database as a subpage in the specified parent page, with the specified properties schema. Currently, the parent of a new database must be a Notion page or a wiki database.
//	@Tags			notion
//	@Accept			json
//	@Produce		json
//	@Param			pageId	path	string	true	"Page ID"
//	@Param 			request body 	requestModel.NotionCreateDatabaseRequest true	"Request Body"
//	@Success		200		{array}		responseModel.NotionCreateDatabaseResponse
//	@Failure		400		{string}	string			"Invalid input"
//	@Router			/api/v1/notion/createDatabase/{pageId} [post]
func (c *Controller) CreateNotionDatabase(ctx *gin.Context) {
	// Get Authorization from config
	configHandler := util.NewConfigHandler()
	auth := configHandler.GetSecretConfig().Get("Authorization")

	pageId := ctx.Param("pageId")
	// Bind the request body to struct
	var requests requestModel.NotionCreateDatabaseRequest
	if err := ctx.ShouldBindJSON(&requests); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Marshal the struct to json
	titleJson, err := json.Marshal(requests.Title)
	if err != nil {
		log.Fatalln(err)
	}
	propertiesJson, err := json.Marshal(requests.Properties)
	if err != nil {
		log.Fatalln(err)
	}

	// Send the request to Notion API
	client := handler.NewClient()
	header := map[string]string{
		"Authorization":  auth.(string),
		"Notion-Version": "2022-06-28",
		"Content-Type":   "application/json",
	}
	bodyString := `{
		"is_inline": true,
		"parent": {
			"type": "page_id",
			"page_id": "` + pageId + `"
		},
		"title": ` + string(titleJson) + `,
		"properties": ` + string(propertiesJson) + ` 
	}`
	body := []byte(bodyString)

	response, err := client.Post("https://api.notion.com/v1/databases", header, body)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	// Change the response body to []byte type
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	bodyStr := string(responseBody)
	var data []byte = []byte(bodyStr)

	// Unmarshal the response body to struct
	var responseCreateNotionDatabase responseModel.NotionCreateDatabaseResponse
	json.Unmarshal(data, &responseCreateNotionDatabase)

	ctx.JSON(http.StatusOK, responseCreateNotionDatabase)
}

// QueryNotionDatabase godoc
//
//	@Summary		Query Notion Database
//	@Description	Queries a database, returning a paginated array of `Page` objects within the database.
//	@Tags			notion
//	@Accept			json
//	@Produce		json
//	@Param			databaseId	path	string	true	"Database ID"
//	@Param 			request body 	requestModel.NotionQueryDatabaseRequest true	"Request Body"
//	@Success		200		{array}		responseModel.NotionQueryDatabaseResponse
//	@Failure		400		{string}	string			"Invalid input"
//	@Router			/api/v1/notion/queryDatabase/{databaseId} [post]
func (c *Controller) QueryNotionDatabase(ctx *gin.Context) {
	databaseId := ctx.Param("databaseId")

	var requests requestModel.NotionQueryDatabaseRequest
	if err := ctx.ShouldBindJSON(&requests); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := infra.QueryNotionDatabaseService(databaseId, requests)

	ctx.JSON(http.StatusOK, response)
}

// Create a Page godoc
//
//	@Summary		Create a new Notion Page
//	@Description	Creates a new page in the specified database or as a child of an existing page.
//	@Tags			notion
//	@Accept			json
//	@Produce		json
//	@Param			databaseId	path	string	true	"Database ID"
//	@Param 			request body 	requestModel.NotionCreateDBPageRequest true	"Request Body"
//	@Success		200		{array}		responseModel.Database
//	@Failure		400		{string}	string			"Invalid input"
//	@Router			/api/v1/notion/createDBPage/{databaseId} [post]
func (c *Controller) CreateNotionDBPage(ctx *gin.Context) {
	databaseId := ctx.Param("databaseId")

	var requests requestModel.NotionCreateDBPageRequest
	if err := ctx.ShouldBindJSON(&requests); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := infra.CreateNotionPageService(databaseId, requests)

	ctx.JSON(http.StatusOK, response)
}
