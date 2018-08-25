package Repositories

import "github.com/jinzhu/gorm"

type Aggregate struct {
	StatsRepository *StatsRepository
	PagesRepository *PagesRepository
}

func NewAggregate(conn *gorm.DB) *Aggregate {
	return &Aggregate{
		StatsRepository: &StatsRepository{conn},
		PagesRepository: &PagesRepository{conn},
	}
}