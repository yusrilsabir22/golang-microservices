package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	MSISDN   string    `gorm:"index:,unique" json:"msisdn"`
	Name     string    `json:"name"`
	Username string    `gorm:"index:,unique" json:"username"`
	Password string    `json:"password"`
}
