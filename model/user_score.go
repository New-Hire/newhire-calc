package model

import (
	"time"
)

type UserScore struct {
	UserId    int64     `gorm:"primarykey:column:user_id;type:bigint;not null"`
	Score     int       `gorm:"column:score;type:int;not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null"`
}

func (UserScore) TableName() string {
	return "user_score"
}
