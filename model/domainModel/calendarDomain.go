package domainModel

import "time"

type CalendarDomain struct {
	Event     string    `json:"event"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"dateTime"`
}
