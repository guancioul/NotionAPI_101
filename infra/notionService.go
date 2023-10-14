package infra

import (
	"encoding/json"
	"io"
	"log"

	"github.com/guancioul/NotionGoogleCalendarIntegration/handler"
	"github.com/guancioul/NotionGoogleCalendarIntegration/model/requestModel"
	"github.com/guancioul/NotionGoogleCalendarIntegration/model/responseModel"
	"github.com/guancioul/NotionGoogleCalendarIntegration/util"
)

func QueryNotionDatabaseService(databaseId string, requests requestModel.NotionQueryDatabaseRequest) responseModel.NotionQueryDatabaseResponse {
	// Get Authorization from config
	configHandler := util.NewConfigHandler()
	auth := configHandler.GetSecretConfig().Get("Authorization")

	// Marshal the struct to json
	requestJson, err := json.Marshal(requests)
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
	body := []byte(requestJson)

	response, err := client.Post("https://api.notion.com/v1/databases/"+databaseId+"/query", header, body)
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
	var responseQueryNotionDatabase responseModel.NotionQueryDatabaseResponse
	json.Unmarshal(data, &responseQueryNotionDatabase)

	return responseQueryNotionDatabase
}

func CreateNotionPageService(databaseId string, requests requestModel.NotionCreateDBPageRequest) responseModel.NotionCreateDatabaseResponse {
	// Get Authorization from config
	configHandler := util.NewConfigHandler()
	auth := configHandler.GetSecretConfig().Get("Authorization")

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
			"parent": {
				"database_id": "` + databaseId + `"
			},
			"properties": ` + string(propertiesJson) + ` 
		}`
	body := []byte(bodyString)

	response, err := client.Post("https://api.notion.com/v1/pages", header, body)
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

	return responseCreateNotionDatabase
}
