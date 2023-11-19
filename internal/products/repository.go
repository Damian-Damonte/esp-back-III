package products

import (
	"context"
	"errors"
	"github.com/aldogayaladh/go-web-1598/internal/domain"
	"github.com/aldogayaladh/go-web-1598/pkg/storage"
)

var (
	ErrEmpty    = errors.New("empty list")
	ErrNotFound = errors.New("product not found")
)

type Repository interface {
	Create(ctx context.Context, producto domain.Producto) (*domain.Producto, error)
	GetAll(ctx context.Context) (*[]domain.Producto, error)
	GetByID(ctx context.Context, id string) (*domain.Producto, error)
	Update(ctx context.Context, producto domain.Producto, id string) (*domain.Producto, error)
	Delete(ctx context.Context, id string) error
}

type repository struct {
	storage storage.Storage
}

func NewMemoryRepository(storage storage.Storage) Repository {
	return &repository{storage: storage}
}

// Create ....
func (r *repository) Create(ctx context.Context, producto domain.Producto) (*domain.Producto, error) {
	product, err := r.storage.Create(ctx, producto)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// GetAll...
func (r *repository) GetAll(ctx context.Context) (*[]domain.Producto, error) {
	products, err := r.storage.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

// GetByID .....
func (r *repository) GetByID(ctx context.Context, id string) (*domain.Producto, error) {
	result, err := r.storage.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Update ...
func (r *repository) Update(
	ctx context.Context,
	producto domain.Producto,
	id string) (*domain.Producto, error) {
	result, err := r.storage.Update(ctx, producto, id)
	if err != nil {
		return nil, err
	}

	return result, nil	
}

// Delete ...
func (r *repository) Delete(ctx context.Context, id string) error {
	err := r.storage.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
