package models

type AccountData struct {
	ID            string         `json:"" gorm:"primary_key"`
	Name          string         `json:"name"`
	AccountEvents []AccountEvent `json:"events" gorm:"ForeignKey:AccountID"`
}
type AccountEvent struct {
	ID        string `json:"" gorm:"primary_key"`
	AccountID string `json:"-" gorm:"index"`
	EventName string `json:"eventName"`
	Created   string `json:"created"`
}
