package database

import (
	"restaurant/models"
)

type IMenuRepository interface {
	Insert(menu models.Menu) error
	FindAll() (error, []models.Menu)
	FindByDate(string) (error, models.Menu)
}
