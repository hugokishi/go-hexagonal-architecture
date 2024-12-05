package ports

import (
	"context"

	"github.com/hugokishi/hexagonal-go/internal/core/models"
)

type ProductRepository interface {
	Save(ctx context.Context, creation models.Product) error
}
