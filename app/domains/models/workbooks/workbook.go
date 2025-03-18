package workbooks

type Workbook struct {
	id          WorkbookId
	description Description
	title       Title
}

func NewUser(id WorkbookId, description Description, title Title) *Workbook {
	return &Workbook{
		id:          id,
		description: description,
		title:       title,
	}
}

func (w *Workbook) Id() int {
	return w.id.Value()
}

func (w *Workbook) Description() string {
	return w.description.Value()
}

func (w *Workbook) Title() string {
	return w.title.Value()
}
