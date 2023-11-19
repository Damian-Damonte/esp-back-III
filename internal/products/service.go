package products

import (
	"context"
	"log"

	"github.com/aldogayaladh/go-web-1598/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) (*[]domain.Producto, error)
	GetByID(ctx context.Context, id string) (*domain.Producto, error)
	Create(ctx context.Context, producto domain.Producto) (*domain.Producto, error)
	Update(ctx context.Context, producto domain.Producto, id string) (*domain.Producto, error)
	Delete(ctx context.Context, id string) error
}

type service struct {
	repository Repository
}

func NewServiceProduct(repository Repository) Service {
	return &service{repository: repository}
}

// GetAll ...
func (s *service) GetAll(ctx context.Context) (*[]domain.Producto, error) {
	listProducts, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("[ProductsService][GetAll] error getting all products", err)
		return nil, err
	}

	return listProducts, nil
}

// GetByID ....
func (s *service) GetByID(ctx context.Context, id string) (*domain.Producto, error) {
	producto, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[ProductsService][GetByID] error getting product by ID", err)
		return nil, err
	}

	return producto, nil
}

// Create ....
func (s *service) Create(ctx context.Context, producto domain.Producto) (*domain.Producto, error) {
	product, err := s.repository.Create(ctx, producto)
	if err != nil {
		log.Println("[ProductsService][Create] error creating product", err)
		return nil, err
	}

	return product, nil
}

// Update ...
func (s *service) Update(ctx context.Context, producto domain.Producto, id string) (*domain.Producto, error) {
	product, err := s.repository.Update(ctx, producto, id)
	if err != nil {
		log.Println("[ProductsService][Update] error updating product by ID", err)
		return nil, err
	}

	return product, nil
}

// Delete ...
func (s *service) Delete(ctx context.Context, id string) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[ProductsService][Delete] error deleting product by ID", err)
		return err
	}

	return nil
}
