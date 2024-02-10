package model

type User struct {
	BaseModel
	Username string `gorm:"column:username;type:varchar(255);not null" json:"username"`
	Password string `gorm:"column:password;type:varchar(255);not null" json:"-"`
	QQNumber string `gorm:"column:qq_number;type:varchar(255);not null" json:"qq_number"`
	QQName   string `gorm:"column:qq_name;type:varchar(255);not null" json:"qq_name"`
	Age      int    `gorm:"column:age;type:int;not null" json:"age"`
	Area     string `gorm:"column:area;type:varchar(255);not null" json:"area"`
	Gender   string `gorm:"column:gender;type:varchar(255);not null" json:"gender"`
}
