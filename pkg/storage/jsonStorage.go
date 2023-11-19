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
	ErrIdTaken = errors.New("ya existe un producto con el id seleccionado")
)

type Storage interface {
	Inicializacion()
	GetAll(ctx context.Context)([]domain.Producto, error)
	GetByID(ctx context.Context, id string) (domain.Producto, error)
	Create(ctx context.Context, producto domain.Producto) (domain.Producto, error)
	Update(ctx context.Context, producto domain.Producto,id string) (domain.Producto, error)
	Delete(ctx context.Context, id string) error
}

type jsonStorage struct {
	Storage []domain.Producto
}

func NewStorage() Storage {
	return &jsonStorage{}
}

func (s *jsonStorage) Inicializacion(){
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

func (s *jsonStorage) GetAll(ctx context.Context, ) ([]domain.Producto, error) {
	if len(s.Storage) == 0 {
		return []domain.Producto{}, ErrEmpty
	}

	return s.Storage, nil
}

func (s *jsonStorage) GetByID(ctx context.Context, id string) (domain.Producto, error) {
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

func (s *jsonStorage) Update(
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

	err := s.UpdateJsonFile()
	if err != nil {
		return domain.Producto{}, err
	}

	return result, nil
}

func (s *jsonStorage) Delete(ctx context.Context, id string) error {
	var result domain.Producto
	for key, value := range s.Storage{
		if value.Id == id {
			result = s.Storage[key]
			s.Storage = append(s.Storage[:key], s.Storage[key+1:]...)
			break
		}
	}

	if result.Id == "" {
		return ErrNotFound
	}

	err := s.UpdateJsonFile()
	if err != nil {
		return err
	}

	return nil
}

func (s *jsonStorage) Create(ctx context.Context, producto domain.Producto) (domain.Producto, error) {
	_, err := s.GetByID(ctx, producto.Id)
	if err == nil {
		return domain.Producto{}, ErrIdTaken
	}

	s.Storage = append(s.Storage, producto)

	err = s.UpdateJsonFile()
	if err != nil {
		return domain.Producto{}, err
	}

	return producto, nil
}


func (s *jsonStorage) UpdateJsonFile() error{
	productosJson, err := json.Marshal(s.Storage)
	if err != nil {
		return ErrMarshal
	}

	err = os.WriteFile(jsonFile, productosJson, 0644)

	if err != nil {
		return ErrWriteFile
	}

	return nil
}