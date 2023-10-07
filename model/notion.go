package model

type NotionCreateDatabaseResponse struct {
	Object         string `json:"object"`
	ID             string `json:"id"`
	CreatedTime    string `json:"created_time"`
	LastEditedTime string `json:"last_edited_time"`
	Icon           struct {
		Type  string `json:"type"`
		Emoji string `json:"emoji"`
	} `json:"icon"`
	Cover struct {
		Type     string `json:"type"`
		External struct {
			URL string `json:"url"`
		} `json:"external"`
	} `json:"cover"`
	URL   string `json:"url"`
	Title []struct {
		Type string `json:"type"`
		Text struct {
			Content string      `json:"content"`
			Link    interface{} `json:"link"`
		} `json:"text"`
		Annotations struct {
			Bold          bool   `json:"bold"`
			Italic        bool   `json:"italic"`
			Strikethrough bool   `json:"strikethrough"`
			Underline     bool   `json:"underline"`
			Code          bool   `json:"code"`
			Color         string `json:"color"`
		} `json:"annotations"`
		PlainText string      `json:"plain_text"`
		Href      interface{} `json:"href"`
	} `json:"title"`
	Properties struct {
		PlusOne struct {
			ID     string   `json:"id"`
			Name   string   `json:"name"`
			Type   string   `json:"type"`
			People struct{} `json:"people"`
		} `json:"+1"`
		InStock struct {
			ID       string   `json:"id"`
			Name     string   `json:"name"`
			Type     string   `json:"type"`
			Checkbox struct{} `json:"checkbox"`
		} `json:"In stock"`
		Price struct {
			ID     string `json:"id"`
			Name   string `json:"name"`
			Type   string `json:"type"`
			Number struct {
				Format string `json:"format"`
			} `json:"number"`
		} `json:"Price"`
		Description struct {
			ID       string   `json:"id"`
			Name     string   `json:"name"`
			Type     string   `json:"type"`
			RichText struct{} `json:"rich_text"`
		} `json:"Description"`
		LastOrdered struct {
			ID   string   `json:"id"`
			Name string   `json:"name"`
			Type string   `json:"type"`
			Date struct{} `json:"date"`
		} `json:"Last ordered"`
		StoreAvailability struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			MultiSelect struct {
				Options []struct {
					ID    string `json:"id"`
					Name  string `json:"name"`
					Color string `json:"color"`
				} `json:"options"`
			} `json:"multi_select"`
		} `json:"Store availability"`
		Photo struct {
			ID    string   `json:"id"`
			Name  string   `json:"name"`
			Type  string   `json:"type"`
			Files struct{} `json:"files"`
		} `json:"Photo"`
		FoodGroup struct {
			ID     string `json:"id"`
			Name   string `json:"name"`
			Type   string `json:"type"`
			Select struct {
				Options []struct {
					ID    string `json:"id"`
					Name  string `json:"name"`
					Color string `json:"color"`
				} `json:"options"`
			} `json:"select"`
		} `json:"Food group"`
		Name struct {
			ID    string   `json:"id"`
			Name  string   `json:"name"`
			Type  string   `json:"type"`
			Title struct{} `json:"title"`
		} `json:"Name"`
	} `json:"properties"`
	Parent struct {
		Type   string `json:"type"`
		PageID string `json:"page_id"`
	} `json:"parent"`
	Archived bool `json:"archived"`
	IsInline bool `json:"is_inline"`
}
