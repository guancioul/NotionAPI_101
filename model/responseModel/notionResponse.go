package responseModel

type NotionCreateDatabaseResponse struct {
	Object         string                   `json:"object"`
	ID             string                   `json:"id"`
	CreatedTime    string                   `json:"created_time"`
	LastEditedTime string                   `json:"last_edited_time"`
	URL            string                   `json:"url"`
	Title          []map[string]interface{} `json:"title"`
	Properties     map[string]interface{}   `json:"properties"`
	Parent         struct {
		Type   string `json:"type"`
		PageID string `json:"page_id"`
	} `json:"parent"`
	Archived bool `json:"archived"`
	IsInline bool `json:"is_inline"`
}

type Database struct {
	Object         string `json:"object"`
	ID             string `json:"id"`
	CreatedTime    string `json:"created_time"`
	LastEditedTime string `json:"last_edited_time"`
	CreatedBy      struct {
		Object string `json:"object"`
		ID     string `json:"id"`
	} `json:"created_by"`
	LastEditedBy struct {
		Object string `json:"object"`
		ID     string `json:"id"`
	} `json:"last_edited_by"`
	Cover  interface{} `json:"cover"`
	Icon   interface{} `json:"icon"`
	Parent struct {
		Type       string `json:"type"`
		DatabaseID string `json:"database_id"`
	} `json:"parent"`
	Archived   bool                   `json:"archived"`
	Properties map[string]interface{} `json:"properties"`
	URL        string                 `json:"url"`
	PublicURL  interface{}            `json:"public_url"`
}

type NotionQueryDatabaseResponse struct {
	Object          string      `json:"object"`
	Results         []Database  `json:"results"`
	NextCursor      interface{} `json:"next_cursor"`
	HasMore         bool        `json:"has_more"`
	Type            string      `json:"type"`
	PageOrDB        struct{}    `json:"page_or_database"`
	DeveloperSurvey string      `json:"developer_survey"`
}
