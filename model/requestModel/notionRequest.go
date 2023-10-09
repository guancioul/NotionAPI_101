package requestModel

type NotionCreateDatabaseRequest struct {
	Title      []map[string]interface{} `json:"title"`
	Properties map[string]interface{}   `json:"properties"`
}

type NotionQueryDatabaseRequest struct {
	Sorts  []map[string]interface{} `json:"sorts"`
	Filter map[string]interface{}   `json:"filter"`
}
