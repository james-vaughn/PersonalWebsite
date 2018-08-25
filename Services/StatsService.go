package Services

import (
	"github.com/james-vaughn/PersonalWebsite/Models"
	"github.com/james-vaughn/PersonalWebsite/Repositories"
)

type StatsService struct {
	StatsRepository *Repositories.StatsRepository
}

func NewStatsService(StatsRepository *Repositories.StatsRepository) *StatsService {
	return &StatsService{StatsRepository: StatsRepository}
}

func (s *StatsService) Create(stat *Models.Stat) {
	s.StatsRepository.Create(stat)
}