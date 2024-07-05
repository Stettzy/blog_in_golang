package comment

type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Likes   int    `json:"likes"`
}

func NewComment() *Comment {
	c := &Comment{}
	return c
}
