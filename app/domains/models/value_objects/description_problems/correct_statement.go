package description_problems

import validation "github.com/go-ozzo/ozzo-validation"

type CorrectStatement struct {
	value string
}

func NewCorrectStatement(value string) (CorrectStatement, error) {
	err := validation.Validate(value,
		validation.Required,
		validation.Length(1, 100),
	)

	if err != nil {
		return CorrectStatement{value: ""}, err
	}

	return CorrectStatement{value: value}, nil
}

func (s *CorrectStatement) Value() string {
	return s.value
}
