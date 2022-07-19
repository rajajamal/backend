package services

import (
	"github.com/rajajamal/resto-backend/models"
	"github.com/rajajamal/resto-backend/repositories"
)

type Menu struct {
	Repository repositories.Menu
}

func (s Menu) GetAll() ([]models.Menu, error) {
	return s.Repository.FindAll()
}

func (s Menu) Get(id int) (models.Menu, error) {
	return s.Repository.Find(id)
}

func (s Menu) Save(menu models.Menu) (models.Menu, error) {
	return menu, s.Repository.Saves(&menu)
}
