package service

import (
	"context"
	"log"

	"forum/internal/entity"
	"forum/internal/repository"
)

type categoryService struct {
	catRepo repository.CategoryRepo
}

func NewCategoryService(catRepo repository.CategoryRepo) CategoryService {
	log.Println("| | category service is done!")
	return &categoryService{
		catRepo: catRepo,
	}
}

func (c *categoryService) CreateCategory(ctx context.Context, categories []entity.Category) ([]int64, error) {
	return c.catRepo.CreateCategory(ctx, categories)
}

func (c *categoryService) GetAllCategories(ctx context.Context) ([]entity.Category, error) {
	return c.catRepo.GetAllCategories(ctx)
}

func (c *categoryService) GetCategoryByID(ctx context.Context, categoryID uint64) (entity.Category, error) {
	return c.catRepo.GetCategoryByID(ctx, categoryID)
}
