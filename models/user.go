package models

// User represents the model for an user
type User struct {
	UserID int8   `json:"user_id" example:"123456"`
	Name   string `json:"name" example:"HidayatAjiSantoso"`
}
