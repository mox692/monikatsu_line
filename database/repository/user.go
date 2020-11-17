package repository

type User struct {
	ID           string
	LineID       string
	UserName     string
	UserPassword string
}

func SelectByID(userID string) *User {
	return &User{}
}

func SelectByUserName(userName string) (*User, error) {
	return &User{}, nil
}
