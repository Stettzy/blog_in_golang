package post

func (p *Post) CreatePost() (int, error) {
	return 0, nil
}

func (p *Post) UpdatePost() (int, error) {
	return 0, nil
}

func (p *Post) RemovePost() (int, error) {
	return 0, nil
}

const createPost = `INSERT INTO posts (title, content) VALUES (?, ?)`
const updatePost = `UPDATE posts (title, content) VALUES (?, ?) WHERE id = ?`
const removePost = `REMOVE FROM posts WHERE id = ?`
