package answer_truths

type Truth struct {
	value bool
}

func NewTruth(value bool) Truth {
	return Truth{value: value}
}

func (t *Truth) Value() bool {
	return t.value
}
