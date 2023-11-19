package sql

import (
	"context"
	"database/sql"

	"github.com/aldogayaladh/go-web-1598/internal/domain"
)

type SqlStore struct {
	DB *sql.DB
}

func (s *SqlStore) GetByID(ctx context.Context, id string) (*domain.Producto, error) {
	var producto domain.Producto

	query := "SELECT * FROM products WHERE id = ?"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&producto.Id, &producto.Name, &producto.Quantity, &producto.CodeValue, &producto.IsPublished, &producto.Expiration, &producto.Price)
	if err != nil {
		return nil, err
	}

	return &producto, nil
}

