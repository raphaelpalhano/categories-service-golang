package use_cases

import (
	"log"

	"github.com/raphaelpalhano/categories-ms/cmd/internal/entities"
	"github.com/raphaelpalhano/categories-ms/cmd/internal/repositories"
)

type createCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewCreateCategoryUseCase(repository repositories.ICategoryRepository) *createCategoryUseCase {
	return &createCategoryUseCase{repository}
}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}
	log.Print(category)

	err = u.repository.Save(category)

	if err != nil {
		return err
	}

	return nil

}
