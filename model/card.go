package model

type Card struct {
	BaseModel
	From    string `gorm:"column:from;type:varchar(255);not null" json:"from"`
	To      string `gorm:"column:to;type:varchar(255);not null" json:"to"`
	Content string `gorm:"column:content;type:longtext;not null" json:"content,omitempty"`
	Key     string `gorm:"column:key;type:varchar(255);not null" json:"key"`
	UserId  int    `gorm:"column:user_id;type:int;not null" json:"user_id"`
}
