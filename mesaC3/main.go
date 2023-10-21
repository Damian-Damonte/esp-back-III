package main

import (
	"fmt"
)

type Product struct {
	ID          int
	Name        string
	Price       float32
	Description string
	Category    string
}

func (p *Product) Save() {
	Products = append(Products, *p)
}

func (p Product) GetAll() {
	for _,v := range Products {
		fmt.Println(v)
	}
}

func getById(id int) (Product, error) {
	for _,v := range Products {
		if v.ID == id {
			return v, nil
		}
	}

	return Product{}, fmt.Errorf("no se encontró el producto con id: %d", id)
}

var Products = []Product{
	{1, "Celular", 120000.00, "Celular Samsung", "Celulares"},
	{2, "Teclado", 25000.00, "Teclado HyperX", "Teclados"},
	{3, "TV LG", 340000.00, "TV LG 43'", "TV"},
}

func main() {
	product4 := Product{4, "Reloj Casio", 25000.00, "Reloj Casio clasico", "Reloj"}

	product4.Save()

	product4.GetAll()

	productById, err := getById(3)
	if err == nil {
		fmt.Println(productById)
	}	else {
		fmt.Println(err.Error())
	}
}