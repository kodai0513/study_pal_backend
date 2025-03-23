package workbooks

type IsPublic struct {
	value bool
}

func NewIsPublic(value bool) IsPublic {
	return IsPublic{value: value}
}

func (i *IsPublic) Value() bool {
	return i.value
}
