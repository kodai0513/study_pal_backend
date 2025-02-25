package app_types

type PageResponse struct {
	PageSize      int    `json:"page_size"`
	PrevPageToken string `json:"prev_page_token"`
	NextPageToken string `json:"next_page_token"`
}

func NewPageResponse(page *Page) *PageResponse {
	return &PageResponse{
		PageSize:      page.PageSize(),
		PrevPageToken: page.PrevPageToken(),
		NextPageToken: page.NextPageToken(),
	}
}
