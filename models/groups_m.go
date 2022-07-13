package models

import (
	"encoding/json"
	"time"
)

type Groups struct {
	Id           int64            `json:"id" orm:"auto" gorm:"primary_key"`
	Name         string           `json:"name" orm:"size(128)"`
	Type         string           `json:"type" orm:"size(64)" default:"PERSONAL"`
	Metadata     *json.RawMessage `json:"metadata" default:"{}"`
	CollectionId int64            `json:"collection_id"`
	CreatedDate  time.Time        `json:"created_date"`
}
