package app_types

type PageResponse struct {
	PageSize   int    `json:"page_size"`
	PrevPageId string `json:"prev_page_id"`
	NextPageId string `json:"next_page_id"`
}

func NewPageResponse(page *Page) *PageResponse {
	return &PageResponse{
		PageSize:   page.PageSize(),
		PrevPageId: page.PrevPageId(),
		NextPageId: page.NextPageId(),
	}
}
