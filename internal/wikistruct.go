package wikistruct

type SearchResult struct {
	Query Query `json:"query"`
}

type Query struct {
	Search []Search `json:"search"`
}

type Search struct {
	Title string `json:"title"`
}

type Page struct {
	Extract     string      `json:"extract"`
	ContentUrls ContentUrls `json:"content_urls"`
}

type ContentUrls struct {
	Desktop Desktop `json:"desktop"`
}

type Desktop struct {
	Page string `json:"page"`
}
