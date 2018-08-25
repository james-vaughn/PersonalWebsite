package Repositories

import "github.com/james-vaughn/PersonalWebsite/Models"

type PagesRepository Repository

func (r *PagesRepository) GetAll(stats *Models.Stat) []Models.Page{
	var pages []Models.Page
	r.DbConn.Find(&pages)

	return pages
}