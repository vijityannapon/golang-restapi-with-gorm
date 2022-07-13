package models

import (
	"encoding/json"
	"time"
)

type Collections struct {
	Id          int64            `json:"id" orm:"auto"gorm:"primary_key"`
	Name        string           `json:"name" orm:"size(128)"`
	Type        string           `json:"type" orm:"size(64)" default:"PERSONAL"`
	Metadata    *json.RawMessage `json:"metadata" default:"{}"`
	Groups      []Groups         `gorm:"foreignKey:CollectionId"`
	CreatedDate time.Time        `json:"created_date"`
}
