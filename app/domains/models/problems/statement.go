package problems

import validation "github.com/go-ozzo/ozzo-validation"

type Statement struct {
	value string
}

func NewStatement(value string) (Statement, error) {
	err := validation.Validate(value,
		validation.Required,
		validation.Length(1, 1000),
	)

	if err != nil {
		return Statement{value: ""}, err
	}

	return Statement{value: value}, nil
}

func (s *Statement) Value() string {
	return s.value
}
