package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/handler"
	"github.com/swaggo/swag/example/celler/model"
	_ "github.com/swaggo/swag/example/celler/model"
	"github.com/swaggo/swag/example/celler/util"
)

// CreateNotionDatabase godoc
//
//	@Summary		Create a new Notion Database
//	@Description	Creates a database as a subpage in the specified parent page, with the specified properties schema. Currently, the parent of a new database must be a Notion page or a wiki database.
//	@Tags			notion
//	@Accept			json
//	@Produce		json
//	@Success		200		{array}		model.NotionCreateDatabaseResponse
//	@Failure		400		{string}	string			"fail"
//	@Router			/api/v1/notion/createDatabase/{pageId} [post]
func (c *Controller) CreateNotionDatabase(ctx *gin.Context) {
	pageId := ctx.Param("pageId")
	configHandler := util.NewConfigHandler()
	auth := configHandler.GetSecretConfig().Get("Authorization")

	client := handler.NewClient()
	header := map[string]string{
		"Authorization":  auth.(string),
		"Notion-Version": "2022-06-28",
		"Content-Type":   "application/json",
	}
	body := []byte(`{
		"is_inline": true,
		"parent": {
			"type": "page_id",
			"page_id": "` + pageId + `"
		},
		"icon": {
			"type": "emoji",
				"emoji": "📝"
		  },
		  "cover": {
			  "type": "external",
			"external": {
				"url": "https://website.domain/images/image.png"
			}
		  },
		"title": [
			{
				"type": "text",
				"text": {
					"content": "Grocery List",
					"link": null
				}
			}
		],
		"properties": {
			"Name": {
				"title": {}
			},
			"Description": {
				"rich_text": {}
			},
			"In stock": {
				"checkbox": {}
			},
			"Food group": {
				"select": {
					"options": [
						{
							"name": "🥦Vegetable",
							"color": "green"
						},
						{
							"name": "🍎Fruit",
							"color": "red"
						},
						{
							"name": "💪Protein",
							"color": "yellow"
						}
					]
				}
			},
			"Price": {
				"number": {
					"format": "dollar"
				}
			},
			"Last ordered": {
				"date": {}
			},
			"Store availability": {
				"type": "multi_select",
				"multi_select": {
					"options": [
						{
							"name": "Duc Loi Market",
							"color": "blue"
						},
						{
							"name": "Rainbow Grocery",
							"color": "gray"
						},
						{
							"name": "Nijiya Market",
							"color": "purple"
						},
						{
							"name": "Gus's Community Market",
							"color": "yellow"
						}
					]
				}
			},
			"+1": {
				"people": {}
			},
			"Photo": {
				"files": {}
			}
		}
	}`)
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
	var responseCreateNotionDatabase model.NotionCreateDatabaseResponse
	json.Unmarshal(data, &responseCreateNotionDatabase)

	ctx.JSON(http.StatusOK, responseCreateNotionDatabase)
}
