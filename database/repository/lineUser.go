package repository

import (
	"github.com/mox692/monikatsu_line/database"

	"golang.org/x/xerrors"
)

type LineUser struct {
	UserID        string
	LineUserToken string
}

func (lu *LineUser) SelectLineUserByID(userID string) *LineUser {
	return &LineUser{}
}

func (lu *LineUser) InsertLineUser() error {
	stmt, err := database.Conn.Prepare("INSERT INTO line_user (user_id, line_user_token) VALUES (?, ?);")
	if err != nil {
		return xerrors.Errorf("db.Conn.Prepare err : %w", err)
	}
	_, err = stmt.Exec(lu.UserID, lu.LineUserToken)
	if err != nil {
		return xerrors.Errorf("stmt.Exec err : %w", err)
	}
	return nil
}
