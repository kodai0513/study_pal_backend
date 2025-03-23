package workbook_categories

import validation "github.com/go-ozzo/ozzo-validation"

type Name struct {
	value string
}

func NewName(value string) (Name, error) {
	err := validation.Validate(value,
		validation.Required,
		validation.Length(1, 30),
	)

	if err != nil {
		return Name{value: ""}, err
	}

	return Name{value: value}, nil
}

func (n *Name) Value() string {
	return n.value
}
