package post

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewPost() *Post {
	p := &Post{}
	return p
}
