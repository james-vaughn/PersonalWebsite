package Services

import (
	"fmt"
	"github.com/james-vaughn/PersonalWebsite/Models"
	"github.com/james-vaughn/PersonalWebsite/Repositories"
)

type PagesService struct {
	PagesRepository *Repositories.PagesRepository
}

func NewPagesService(PagesRepository *Repositories.PagesRepository) *PagesService {
	return &PagesService{PagesRepository: PagesRepository}
}

func (s *PagesService) GetPagesFor(controller string) []Models.Page{
	return s.PagesRepository.GetPagesFor(controller)
}

func (s *PagesService) GetPage(pages []Models.Page, id uint) Models.Page {
	for _, page := range pages {
		if page.ID == id {
			return page
		}
	}

	return Models.Page{}
}

func (s *PagesService) GetUrlFor(page Models.Page) string {
	return fmt.Sprintf("/%s/%s", page.Controller, page.Url)
}