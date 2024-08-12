package repositories

import "github.com/raphaelpalhano/categories-ms/cmd/internal/entities"

type ICategoryRepository interface {
	Save(category *entities.Category) error
	List() ([]*entities.Category, error)
}
