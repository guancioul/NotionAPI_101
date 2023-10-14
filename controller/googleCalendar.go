package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guancioul/NotionGoogleCalendarIntegration/infra"
)

// GetGoogleCalendarEventList godoc
//
//	@Summary		Get Google Calendar Event List
//	@Description	Get a list of events from Google Calendar
//	@Tags			googleCalendar
//	@Accept			json
//	@Produce		json
//	@Param			calendarId	path	string	true	"Database ID"
//	@Param			timeMin	query	string	false	"timeMin"
//	@Param			timeMax query	string	false	"timeMax"
//	@Success		200		{array}		responseModel.CalendarEvents "success"
//	@Failure		400		{string}	string			"fail"
//	@Failure		404		{string}	string			"Calendar not found"
//	@Router			/api/v1/googleCalendar/getEventList/{calendarId} [get]
func (c *Controller) GetGoogleCalendarEventList(ctx *gin.Context) {
	calendarId := ctx.Param("calendarId")

	timeMin := ctx.Request.URL.Query().Get("timeMin")
	err := infra.CheckTimeFormat(timeMin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error parsing timeMin:": err.Error()})
	}

	timeMax := ctx.Request.URL.Query().Get("timeMax")
	err = infra.CheckTimeFormat(timeMax)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error parsing timeMax:": err.Error()})
	}

	eventList := infra.GetGoogleCalendarEventListService(calendarId, timeMin, timeMax, ctx)

	ctx.JSON(http.StatusOK, eventList)
}
