package repository

type User struct {
	ID           string
	LineID       string
	UserName     string
	UserPassword string
}

func SelectUserByID(userID string) *User {
	return &User{}
}

func SelectUserByUserName(userName string) (*User, error) {
	return &User{}, nil
}

func SelectUserByPassword(pass string) (*User, error) {
	return &User{}, nil
}
