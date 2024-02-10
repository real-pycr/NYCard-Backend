package model

type Poem struct {
	BaseModel
	Title         string `gorm:"column:title;type:varchar(255);not null" json:"title"`
	ContentFirst  string `gorm:"column:content_first;type:varchar(255);not null" json:"content_first"`
	ContentSecond string `gorm:"column:content_second;type:varchar(255);not null" json:"content_second"`
	Author        string `gorm:"column:author;type:varchar(255);not null" json:"author"`
}
