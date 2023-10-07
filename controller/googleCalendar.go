package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/util"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// GetGoogleCalendarEventList godoc
//
//	@Summary		Get Google Calendar Event List
//	@Description	Get Google Calendar Event List
//	@Tags			googleCalendar
//	@Accept			json
//	@Produce		json
//	@Success		200		{array}		string "success"
//	@Failure		400		{string}	string			"fail"
//	@Router			/api/v1/googleCalendar/getEventList [get]
func (c *Controller) GetGoogleCalendarEventList(ctx *gin.Context) {
	configHandler := util.NewConfigHandler()
	auth := configHandler.GetSecretConfig().Get("GoogleSecretKey")
	calendarService, err := calendar.NewService(ctx, option.WithAPIKey(auth.(string)))
	if err != nil {
		panic(err)
	}
	eventList, err := calendarService.Events.List("80b1541526db7f79f0ff886c7b08774f0af842a5799a4a2c364d519e88768925@group.calendar.google.com").Do()
	if err != nil {
		panic(err)
	}
	for _, item := range eventList.Items {
		log.Println(item.Summary)
	}
}
