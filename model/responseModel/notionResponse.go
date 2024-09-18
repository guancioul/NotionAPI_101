package responseModel

import "time"

type User struct {
	Object string `json:"object"`
	ID     string `json:"id"`
}

type NotionPageProperties struct {
	Date struct {
		ID       string `json:"id"`
		Type     string `json:"type"`
		DateInfo struct {
			Start    string  `json:"start"`
			End      string  `json:"end"`
			TimeZone *string `json:"time_zone,omitempty"`
		} `json:"date"`
	} `json:"Date"`
	Title struct {
		ID    string `json:"id"`
		Type  string `json:"type"`
		Title []struct {
			Type        string `json:"type"`
			TextContent struct {
				Content     string  `json:"content"`
				Link        *string `json:"link,omitempty"`
				Annotations struct {
					Bold          bool   `json:"bold"`
					Italic        bool   `json:"italic"`
					Strikethrough bool   `json:"strikethrough"`
					Underline     bool   `json:"underline"`
					Code          bool   `json:"code"`
					Color         string `json:"color"`
				} `json:"annotations"`
				PlainText string  `json:"plain_text"`
				Href      *string `json:"href,omitempty"`
			} `json:"text"`
		} `json:"title"`
	} `json:"Title"`
}

type NotionPage struct {
	Object         string      `json:"object"`
	ID             string      `json:"id"`
	CreatedTime    time.Time   `json:"created_time"`
	LastEditedTime time.Time   `json:"last_edited_time"`
	CreatedBy      User        `json:"created_by"`
	LastEditedBy   User        `json:"last_edited_by"`
	Cover          interface{} `json:"cover"` // This can be any type, so I used interface{}
	Icon           interface{} `json:"icon"`  // This can be any type, so I used interface{}
	Parent         struct {
		Type       string `json:"type"`
		DatabaseID string `json:"database_id"`
	} `json:"parent"`
	Archived   bool                 `json:"archived"`
	Properties NotionPageProperties `json:"properties"`
	URL        string               `json:"url"`
	PublicURL  interface{}          `json:"public_url"` // This can be any type, so I used interface{}
}

type QueryNotionDatabaseResponse struct {
	Object         string       `json:"object"`
	Results        []NotionPage `json:"results"`
	NextCursor     interface{}  `json:"next_cursor"` // This can be any type, so I used interface{}
	HasMore        bool         `json:"has_more"`
	Type           string       `json:"type"`
	PageOrDatabase interface{}  `json:"page_or_database"` // This can be any type, so I used interface{}
	RequestID      string       `json:"request_id"`
}
