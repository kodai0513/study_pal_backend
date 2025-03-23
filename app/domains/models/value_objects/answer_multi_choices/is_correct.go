package answer_multi_choices

type IsCorrect struct {
	value bool
}

func NewIsCorrect(value bool) IsCorrect {
	return IsCorrect{value: value}
}

func (i *IsCorrect) Value() bool {
	return i.value
}
