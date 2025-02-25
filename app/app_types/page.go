package app_types

type Page struct {
	pageSize      int
	prevPageToken string
	nextPageToken string
}

func NewPage(pageSize int, prevPageToken string, nextPageToken string) *Page {
	return &Page{
		pageSize:      pageSize,
		prevPageToken: prevPageToken,
		nextPageToken: nextPageToken,
	}
}

func (p *Page) PageSize() int {
	return p.pageSize
}

func (p *Page) PrevPageToken() string {
	return p.prevPageToken
}

func (p *Page) NextPageToken() string {
	return p.nextPageToken
}
