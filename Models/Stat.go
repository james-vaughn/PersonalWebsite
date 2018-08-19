package Models

import (
	"github.com/jinzhu/gorm"
)

type Stat struct {
	gorm.Model
	IpAddress string
	Page      string
}