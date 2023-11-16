package storage

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/aldogayaladh/go-web-1598/internal/domain"
)

const (
	jsonFile = "data.json"
)

var (
	ErrEmpty    = errors.New("empty list")
	ErrNotFound = errors.New("product not found")
	ErrMarshal = errors.New("error convertir los productos a json")
	ErrWriteFile = errors.New("error al escribir en el archivo")
)

type Storage interface {
	Inicializacion()
	GetByID(ctx context.Context, id string) (domain.Producto, error)
	Update(ctx context.Context, producto domain.Producto,id string) (domain.Producto, error)
}

type storage struct {
	Storage []domain.Producto
}

func NewStorage() Storage {
	return &storage{}
}

func (s *storage) Inicializacion(){
	productosJson, err := os.ReadFile(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var productos []domain.Producto

	err = json.Unmarshal([]byte(productosJson), &productos)

	if err != nil {
		log.Fatal(err)
	}

	s.Storage = productos
}

func (s *storage) GetByID(ctx context.Context, id string) (domain.Producto, error) {
	var result domain.Producto
	for _, value := range s.Storage {
		if value.Id == id {
			result = value
			break
		}
	}

	if result.Id == "" {
		return domain.Producto{}, ErrNotFound
	}

	return result, nil
}

func (s *storage) Update(
	ctx context.Context,
	producto domain.Producto,
	id string) (domain.Producto, error) {

	var result domain.Producto
	for key, value := range s.Storage {
		if value.Id == id {
			producto.Id = id
			s.Storage[key] = producto
			result = s.Storage[key]
			break
		}
	}

	if result.Id == "" {
		return domain.Producto{}, ErrNotFound
	}

	productosJson, err := json.Marshal(s.Storage)
	if err != nil {
		return domain.Producto{}, ErrMarshal
	}

	err = os.WriteFile(jsonFile, productosJson, 0644)

	if err != nil {
		return domain.Producto{}, ErrWriteFile
	}

	return result, nil
}