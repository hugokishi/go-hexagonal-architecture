package repository

import (
	"context"

	"github.com/hugokishi/hexagonal-go/internal/core/drivers/db"
	"github.com/hugokishi/hexagonal-go/internal/core/models"
)

func (u *DB) Save(ctx context.Context, creation models.Product) error {
	if tx := u.cnn.WithContext(ctx).Save(creation); db.IsError(tx.Error) {
		return tx.Error
	}
	return nil
}
