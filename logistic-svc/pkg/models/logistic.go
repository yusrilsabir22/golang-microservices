package models

import "github.com/google/uuid"

type Logistic struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	LogisticName    string    `json:"logistic_name"`
	Amount          int       `json:"amount"`
	DestinationName string    `json:"destination_name"`
	OriginName      string    `json:"origin_name"`
	Duration        string    `json:"duration"`
}
