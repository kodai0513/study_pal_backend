package app_types

type Page struct {
	pageSize   int
	prevPageId string
	nextPageId string
}

func NewPage(pageSize int, prevPageId string, nextPageId string) *Page {
	return &Page{
		pageSize:   pageSize,
		prevPageId: prevPageId,
		nextPageId: nextPageId,
	}
}

func (p *Page) PageSize() int {
	return p.pageSize
}

func (p *Page) PrevPageId() string {
	return p.prevPageId
}

func (p *Page) NextPageId() string {
	return p.nextPageId
}

func (p *Page) SetPageSize(value int) {
	p.pageSize = value
}

func (p *Page) SetPrevPageId(value string) {
	p.prevPageId = value
}

func (p *Page) SetNextPageId(value string) {
	p.nextPageId = value
}
