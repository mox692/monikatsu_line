package repository

import (
	"database/sql"

	"github.com/mox692/monikatsu_line/database"
	"golang.org/x/xerrors"
)

type WakeupData struct {
	ID         string
	UserID     string
	WakeupDay  string
	WakeupTime string
	CreatedAt  string
	UpdatedAt  string
}

func CreateWakeupData() *WakeupData {
	return &WakeupData{}
}

func (ws *WakeupData) SelectByID(id string) *WakeupData {
	return &WakeupData{}
}

func (ws *WakeupData) SelectByUserID(userID string) (*WakeupData, error) {
	// TODO: orderbyを使うのに、wakeupDayの型を見直した方がいいかも？？
	row := database.Conn.QueryRow("select * from  monikatsu.wakeup_data where user_id = ? orderby wakeupDay limit 1;", userID)
	return convertToWakeupData(row)
}

func (ws *WakeupData) Insert() error {
	stmt, err := database.Conn.Prepare("INSERT INTO wakeup_data (id, user_id, wakeup_day, wakeup_time) VALUES (?, ?, ?, ?);")
	if err != nil {
		return xerrors.Errorf("db.Conn.Prepare err : %w", err)
	}
	_, err = stmt.Exec(ws.ID, ws.UserID, ws.WakeupDay, ws.WakeupTime)
	if err != nil {
		return xerrors.Errorf("stmt.Exec err : %w", err)
	}
	return nil
}

func convertToWakeupData(row *sql.Row) (*WakeupData, error) {
	wd := WakeupData{}
	err := row.Scan(&wd.ID, &wd.UserID, &wd.WakeupDay, &wd.WakeupTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, xerrors.Errorf("row.Scan error: %w", err)
	}
	return &wd, nil
}
