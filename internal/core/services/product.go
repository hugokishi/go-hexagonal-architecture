package services

import (
	"context"

	"github.com/hugokishi/hexagonal-go/internal/core/models"
	"github.com/hugokishi/hexagonal-go/internal/core/ports"
)

type ProductService struct {
	repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) *ProductService {
	return &ProductService{repo}
}

func (u *ProductService) Save(ctx context.Context, creation models.Product) error {
	return u.repo.Save(ctx, creation)
}
