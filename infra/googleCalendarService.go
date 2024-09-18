package infra

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/guancioul/NotionGoogleCalendarIntegration/model/responseModel"
	"github.com/guancioul/NotionGoogleCalendarIntegration/util"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func GetGoogleCalendarEventListService(calendarId string, timeMin string, timeMax string, ctx *gin.Context) *calendar.Events {
	configHandler := util.NewConfigHandler()
	auth := configHandler.GetSecretConfig().Get("GoogleSecretKey")

	calendarService, err := calendar.NewService(ctx, option.WithAPIKey(auth.(string)))
	if err != nil {
		panic(err)
	}

	eventList := &calendar.Events{}
	switch {
	case timeMin != "" && timeMax != "":
		eventList, err = calendarService.Events.List(calendarId).TimeMin(timeMin).TimeMax(timeMax).Do()
	case timeMax != "":
		eventList, err = calendarService.Events.List(calendarId).TimeMax(timeMax).Do()
	case timeMin != "":
		eventList, err = calendarService.Events.List(calendarId).TimeMin(timeMin).Do()
	default:
		eventList, err = calendarService.Events.List(calendarId).Do()
	}
	if err != nil {
		panic(err)
	}

	var response responseModel.CalendarEvents
	response.Kind = eventList.Kind
	response.Etag = eventList.Etag
	response.Summary = eventList.Summary
	response.Description = eventList.Description
	response.Updated = eventList.Updated
	response.TimeZone = eventList.TimeZone
	response.AccessRole = eventList.AccessRole
	response.NextSyncToken = eventList.NextSyncToken

	for _, item := range eventList.Items {
		response.Items = append(response.Items, responseModel.Event{
			Kind:     item.Kind,
			Etag:     item.Etag,
			Id:       item.Id,
			Status:   item.Status,
			HtmlLink: item.HtmlLink,
			Created:  item.Created,
			Updated:  item.Updated,
			Summary:  item.Summary,
			Creator: struct {
				Email string `json:"email"`
			}{
				Email: item.Creator.Email,
			},
			Organizer: struct {
				Email       string `json:"email"`
				DisplayName string `json:"displayName"`
				Self        bool   `json:"self"`
			}{
				Email:       item.Organizer.Email,
				DisplayName: item.Organizer.DisplayName,
				Self:        item.Organizer.Self,
			},
			Start: struct {
				DateTime string `json:"dateTime"`
				TimeZone string `json:"timeZone"`
			}{
				DateTime: item.Start.DateTime,
				TimeZone: item.Start.TimeZone,
			},
			End: struct {
				DateTime string `json:"dateTime"`
				TimeZone string `json:"timeZone"`
			}{
				DateTime: item.End.DateTime,
				TimeZone: item.End.TimeZone,
			},
			ICalUID:   item.ICalUID,
			EventType: item.EventType,
		})
	}

	return eventList
}

func CheckTimeFormat(timeStr string) error {
	if timeStr != "" {
		_, err := time.Parse("2006-01-02T15:04:05+00:00", timeStr)
		if err != nil {
			return err
		}
	}
	return nil
}
