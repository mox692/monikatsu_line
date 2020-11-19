package repository

type WakeupData struct {
	ID         string
	UserID     string
	WakeupDay  string
	WakeupTime string
}

func (ws *WakeupData) SelectByID(id string) *WakeupData {
	return &WakeupData{}
}
