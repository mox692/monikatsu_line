package database

type LineUser struct {
	ID        string
	UserID    string
	LineToken string
	LineName  string
}

func (u *LineUser) SelectByID(userID string) *LineUser {
	return &LineUser{}
}
