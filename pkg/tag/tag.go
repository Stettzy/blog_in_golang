package tag

type Tag struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func NewTag() *Tag {
	t := &Tag{}
	return t
}
