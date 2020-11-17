package database

type User struct {
	ID           string
	LineID       string
	UserName     string
	UserPassword string
}

func (u *User) SelectByID(userID string) *User {
	return &User{}
}
