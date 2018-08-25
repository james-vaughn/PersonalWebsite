package Repositories

import "github.com/james-vaughn/PersonalWebsite/Models"

type PagesRepository Repository

func (r *PagesRepository) GetPagesFor(controller string) []Models.Page{
	var pages []Models.Page
	r.DbConn.Where("controller = ?", controller).Find(&pages)

	return pages
}