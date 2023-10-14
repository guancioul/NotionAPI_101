package controller

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/guancioul/NotionGoogleCalendarIntegration/infra"
	"github.com/guancioul/NotionGoogleCalendarIntegration/model/domainModel"
	"github.com/guancioul/NotionGoogleCalendarIntegration/model/requestModel"
)

var formatString = "2006-01-02T15:04:05-07:00"

// SyncGoogleCalendarToNotion godoc
//
//		@Summary		Sync Google Calendar to Notion
//		@Description	Sync Google Calendar to Notion for one day
//		@Tags			sync
//		@Accept			json
//		@Produce		json
//		@Param			databaseId	path	string	true	"Page ID"
//		@Param			calendarId	path	string	true	"Calendar ID"
//		@Success		200		{array}		responseModel.NotionCreateDatabaseResponse
//		@Failure		400		{string}	string			"Invalid input"
//	 	@Router			/api/v1/sync/syncNotionToGoogleCalendar/{databaseId}/{calendarId} [post]
func (c *Controller) SyncGoogleCalendarToNotion(ctx *gin.Context) {
	databaseId := ctx.Param("databaseId")
	calendarId := ctx.Param("calendarId")

	taipeiTimeZone, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		fmt.Println("Error loading time zone:", err)
		return
	}

	// Get the current time in the Asia/Taipei time zone
	currentTimeInTaipei := time.Now().In(taipeiTimeZone)
	startOfDay := time.Date(currentTimeInTaipei.Year(), currentTimeInTaipei.Month(), currentTimeInTaipei.Day(), 0, 0, 0, 0, currentTimeInTaipei.Location())

	notionQueryRequest := requestModel.NotionQueryDatabaseRequest{
		Filter: map[string]interface{}{
			"and": []map[string]interface{}{
				{
					"property": "Date",
					"date": map[string]string{
						"on_or_after": startOfDay.Format(formatString),
					},
				},
				{
					"property": "Date",
					"date": map[string]string{
						"before": currentTimeInTaipei.Format(formatString),
					},
				},
			},
		},
		Sorts: []map[string]interface{}{},
	}

	notionQueryResponse := infra.QueryNotionDatabaseService(databaseId, notionQueryRequest)
	googleCalendarEventResponse := infra.GetGoogleCalendarEventListService(calendarId, startOfDay.Format(formatString), currentTimeInTaipei.Format(formatString), ctx)
	eventMap := make(map[string]domainModel.CalendarDomain)

	for _, notionPage := range notionQueryResponse.Results {
		title, ok := notionPage.Properties["Title"].(map[string]interface{})
		if !ok {
			fmt.Println("Invalid or missing 'Title' property")
			return
		}

		titleData, ok := title["title"].([]interface{})
		if !ok {
			fmt.Println("Invalid or missing 'title' property")
			return
		}

		titleFirst, ok := titleData[0].(map[string]interface{})
		if !ok {
			fmt.Println("Invalid or missing 'title' property")
			return
		}

		titleText, ok := titleFirst["text"].(map[string]interface{})
		if !ok {
			fmt.Println("Invalid or missing 'text' property")
			return
		}

		titleContent, ok := titleText["content"].(string)
		if !ok {
			fmt.Println("Invalid or missing 'content' property")
			return
		}
		if _, ok := eventMap[titleContent]; !ok {
			properties, ok := notionPage.Properties["Date"].(map[string]interface{})
			if !ok {
				fmt.Println("Invalid or missing 'Date' property")
				return
			}

			date, ok := properties["date"].(map[string]interface{})
			if !ok {
				fmt.Println("Invalid or missing 'date' property")
				return
			}

			start, ok := date["start"].(string)
			if !ok {
				fmt.Println("Invalid or missing 'start' property")
				return
			}

			end, ok := date["start"].(string)
			if !ok {
				fmt.Println("Invalid or missing 'start' property")
				return
			}

			startTime, err := time.Parse(time.RFC3339, start)
			if err != nil {
				fmt.Println("Error parsing time:", err)
				return
			}
			endTime, err := time.Parse(time.RFC3339, end)
			if err != nil {
				fmt.Println("Error parsing time:", err)
				return
			}

			eventMap[titleContent] = domainModel.CalendarDomain{
				Event:     titleContent,
				StartTime: startTime,
				EndTime:   endTime,
			}
		}
	}

	for _, event := range googleCalendarEventResponse.Items {
		if _, ok := eventMap[event.Summary]; !ok {
			infra.CreateNotionPageService(databaseId, requestModel.NotionCreateDBPageRequest{
				Properties: map[string]interface{}{
					"Title": []map[string]interface{}{
						{
							"text": map[string]string{
								"content": event.Summary,
							},
						},
					},
					"Date": map[string]interface{}{
						"start": event.Start.DateTime,
						"end":   event.End.DateTime,
					},
				},
			})
		}
	}
}
