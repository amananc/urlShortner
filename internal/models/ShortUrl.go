package models

type URL struct {
	Id          string `json:"id" gorm:"primaryKey"`
	Hash        string `json:"hash"`
	OriginalUrl string `json:"originalUrl"`
	CreatedAt   string `json:"createdAt"`
}
