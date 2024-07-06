package user

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func NewUser() User {
	u := User{}
	return u
}

func (u *User) AssignRegisterData(username, email, password string) {
	u.Username = username
	u.Email = email
	u.Password = password
}
