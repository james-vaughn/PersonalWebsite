package Models

import "github.com/jinzhu/gorm"

type Page struct {
	gorm.Model
	Title 	string
	Url  	string
	PrevId  int
	NextId  int
}