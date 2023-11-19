package pkg

import (
	"context"

	"github.com/aldogayaladh/go-web-1598/internal/domain"
)

type Storage interface {
	Inicializacion()
	GetAll(ctx context.Context) (*[]domain.Producto, error)
	GetByID(ctx context.Context, id string) (*domain.Producto, error)
	Create(ctx context.Context, producto domain.Producto) (*domain.Producto, error)
	Update(ctx context.Context, producto domain.Producto, id string) (*domain.Producto, error)
	Delete(ctx context.Context, id string) error
}