package repository

// 1:1 User
type WakeupScore struct {
	UserID string
	Score  string
}

func (ws *WakeupScore) SelectByID(userID string) *WakeupScore {
	return &WakeupScore{}
}
