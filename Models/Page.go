package Models

import "github.com/jinzhu/gorm"

type Page struct {
	gorm.Model
	Title       string
	Description string
	Controller  string
	Url         string
	PrevId      uint
	NextId      uint
}
