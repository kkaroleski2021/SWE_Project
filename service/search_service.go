package service

//this file contains the struct and functions for the search history

import "go/entity"

type SearchRecord interface {
	Save(entity.Search) entity.Search
	FindAll() []entity.Search
}

type searchRecord struct {
	searches []entity.Search
}

func New() SearchRecord {
	return &searchRecord{}
}

func (service *searchRecord) Save(search entity.Search) entity.Search {
	service.searches = append(service.searches, search)
	return search
}

func (service *searchRecord) FindAll() []entity.Search {
	return service.searches
}
