package articles

type PostId struct {
	value int
}

func NewPostId(value int) PostId {
	return PostId{value: value}
}

func (p *PostId) Value() int {
	return p.value
}
