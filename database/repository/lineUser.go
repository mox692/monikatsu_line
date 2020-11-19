package repository

type LineUser struct {
	ID        string
	UserID    string
	LineToken string
	LineName  string
}

func SelectLineUserByID(userID string) *LineUser {
	return &LineUser{}
}

func InsertLineUser() error {
	// stmt, err := db.Conn.Prepare("INSERT INTO user (id, auth_token, name, high_score, coin) VALUES (?, ?, ?, ?, ?);")
	// if err != nil {
	// 	return xerrors.Errorf("db.Conn.Prepare err : %w", err)
	// }
	// _, err = stmt.Exec(record.ID, record.AuthToken, record.Name, record.HighScore, record.Coin)
	// if err != nil {
	// 	return xerrors.Errorf("stmt.Exec err : %w", err)
	// }
	return nil
}
