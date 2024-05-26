package model

type Tags struct {
	Id   int    `gorm:"type:int;primary_key;AUTO_INCREMENT"`
	Name string `gorm:"type:varchar(255);not null"`
}
