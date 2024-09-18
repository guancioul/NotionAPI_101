package responseModel

type Event struct {
	Kind     string `json:"kind"`
	Etag     string `json:"etag"`
	Id       string `json:"id"`
	Status   string `json:"status"`
	HtmlLink string `json:"htmlLink"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
	Summary  string `json:"summary"`
	Creator  struct {
		Email string `json:"email"`
	} `json:"creator"`
	Organizer struct {
		Email       string `json:"email"`
		DisplayName string `json:"displayName"`
		Self        bool   `json:"self"`
	} `json:"organizer"`
	Start struct {
		DateTime string `json:"dateTime"`
		TimeZone string `json:"timeZone"`
	} `json:"start"`
	End struct {
		DateTime string `json:"dateTime"`
		TimeZone string `json:"timeZone"`
	} `json:"end"`
	ICalUID   string `json:"iCalUID"`
	Sequence  int    `json:"sequence"`
	EventType string `json:"eventType"`
}

type CalendarEvents struct {
	Kind             string     `json:"kind"`
	Etag             string     `json:"etag"`
	Summary          string     `json:"summary"`
	Description      string     `json:"description"`
	Updated          string     `json:"updated"`
	TimeZone         string     `json:"timeZone"`
	AccessRole       string     `json:"accessRole"`
	DefaultReminders []struct{} `json:"defaultReminders"`
	NextSyncToken    string     `json:"nextSyncToken"`
	Items            []Event    `json:"items"`
}
