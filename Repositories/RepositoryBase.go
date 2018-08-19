package Repositories

import (
	"github.com/jinzhu/gorm"
)

type Repository struct {
	DbConn	*gorm.DB
}