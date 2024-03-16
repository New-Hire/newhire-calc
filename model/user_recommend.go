package model

type UserRecommend struct {
	Id             uint  `gorm:"primarykey:column:id;type:bigint;not null"`
	RaterId        uint  `gorm:"column:rater_id;type:bigint;not null"`
	RaterCompanyId uint  `gorm:"column:rater_company_id;type:bigint;not null"`
	UserId         uint  `gorm:"column:user_id;type:bigint;not null"`
	UserCompanyId  uint  `gorm:"column:bigint unsigned;type:bigint"`
	level          uint8 `gorm:"column:level;type:tinyint;not null"`
}

func (UserRecommend) TableName() string {
	return "user_recommend"
}
