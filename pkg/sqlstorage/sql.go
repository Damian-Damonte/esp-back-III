package sqlstorage

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/aldogayaladh/go-web-1598/internal/domain"
	"github.com/aldogayaladh/go-web-1598/pkg"
	"github.com/aldogayaladh/go-web-1598/pkg/timeformater"
	_ "github.com/go-sql-driver/mysql"
)

var (
	ErrNotFound  = errors.New("product not found")
)

type SqlStore struct {
	DB *sql.DB
}

func NewSqlStorage() pkg.Storage {
	return &SqlStore{}
}

func (s *SqlStore) Inicializacion() {
	dataSource := "root:damian@tcp(localhost:3306)/esp_back_3"
	storageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = storageDB.Ping(); err != nil {
		panic(err)
	}
	s.DB = storageDB
}

func (s *SqlStore) GetByID(ctx context.Context, id string) (*domain.Producto, error) {
	var producto domain.Producto
	var expirationString string

	query := "SELECT * FROM products WHERE id = ?"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&producto.Id, &producto.Name, &producto.Quantity, &producto.CodeValue, &producto.IsPublished, &expirationString, &producto.Price)
	if err != nil {
		return nil, err
	}

	expirationTime, err := timeformater.StringToTime(expirationString)
	if err != nil {
		return nil, err
	}
	producto.Expiration = *expirationTime

	return &producto, nil
}

func (s *SqlStore) GetAll(ctx context.Context) (*[]domain.Producto, error) {
	query := "SELECT * FROM products"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productos []domain.Producto

	for rows.Next() {
		var producto domain.Producto
		var expirationString string

		err = rows.Scan(&producto.Id, &producto.Name, &producto.Quantity, &producto.CodeValue, &producto.IsPublished, &expirationString, &producto.Price)
		if err != nil {
			return nil, err
		}

		expirationTime, err := timeformater.StringToTime(expirationString)
		if err != nil {
			return nil, err
		}
		producto.Expiration = *expirationTime

		productos = append(productos, producto)
	}

	return &productos, nil
}

func (s *SqlStore) Create(ctx context.Context, producto domain.Producto) (*domain.Producto, error) {
	stmt, err := s.DB.Prepare("INSERT INTO products(name, quantity, code_value, is_published, expiration, price) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(producto.Name, producto.Quantity, producto.CodeValue,
		producto.IsPublished, timeformater.TimeToString(producto.Expiration), producto.Price)

	if err != nil {
		return nil, err
	}

	insertedId, _ := result.LastInsertId()

	producto.Id = strconv.Itoa(int(insertedId))

	return &producto, nil
}

func (s *SqlStore) Update(ctx context.Context, producto domain.Producto, id string) (*domain.Producto, error) {
	stmt, err := s.DB.Prepare("UPDATE products SET name = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ?, price = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(producto.Name, producto.Quantity, producto.CodeValue, producto.IsPublished,
		timeformater.TimeToString(producto.Expiration), producto.Price, id)

	if err != nil {
		return nil, err
	}

	return &producto, nil
}

func (s *SqlStore) Delete(ctx context.Context, id string) error {
	query := "DELETE FROM products WHERE id = ?"
	res, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}

	cantFilasEliminadas, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if cantFilasEliminadas == 0 {
		return ErrNotFound
	}

	return nil
}
