package service

import (
	model "final-project/pkg/models"
	"final-project/pkg/repository"
)

type ItemService struct {
	repo repository.ItemRepo
}

func NewItemService(repo repository.ItemRepo) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) Create(item model.Item) error {
	if err := s.repo.Create(item); err != nil {
		return err
	}
	return nil
}

func (s *ItemService) GetAll() ([]model.Item, error) {
	items, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *ItemService) GetById(id int) (model.Item, error) {
	item, err := s.repo.GetById(id)
	if err != nil {
		return model.Item{}, err
	}
	return item, nil
}

func (s *ItemService) GetByCategoryId(categoryId int) ([]model.Item, error) {
	items, err := s.repo.GetItemByCategoryId(categoryId)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *ItemService) SortByName(name string) ([]model.Item, error) {
	items, err := s.repo.SortByName(name)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *ItemService) SortByRating(rating string) ([]model.Item, error) {

	// Get the items sorted by rating
	items, err := s.repo.SortByRating(rating)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *ItemService) SortByPrice(price string) ([]model.Item, error) {
	//// Get the items sorted by price
	items, err := s.repo.SortByPrice(price)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *ItemService) Update(item model.Item, id int) error {

	// Update the item
	if err := s.repo.Update(item, id); err != nil {
		return err
	}
	return nil
}

func (s *ItemService) SetRating(rating float32, id int) error {

	// Set the rating for the item
	if err := s.repo.SetRating(rating, id); err != nil {
		return err
	}
	return nil
}

func (s *ItemService) Delete(id int) error {

	// Delete the item
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
