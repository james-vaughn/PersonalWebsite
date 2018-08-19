package Repositories

import (
	"github.com/james-vaughn/PersonalWebsite/Models"
)

type StatsRepository Repository

func (r *StatsRepository) Create(stats *Models.Stat) {
	r.DbConn.Create(stats)
}