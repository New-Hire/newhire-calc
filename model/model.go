package model

type User struct {
	UserId    int64
	CompanyId int64
	Raters    []User
	Score1    int `json:"score1,omitempty"`
	Score2    int `json:"score2,omitempty"`
}

type Node struct {
	RaterId        int64
	UserId         int64
	Deep           int8
	RaterCompanyId int64
	UserCompanyId  int64
	Score1         int
	Score2         int
}
