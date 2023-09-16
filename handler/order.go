package handler

import (
  "fmt"
  "net/http"
)

type Order struct {}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Create an order")
  w.WriteHeader(http.StatusCreated)
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
  fmt.Println("List all orders")
  w.WriteHeader(http.StatusOK)
}

func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {
  fmt.Println("List an order")
  w.WriteHeader(http.StatusOK)
}

func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Update order")
  w.WriteHeader(http.StatusOK)
}

func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Delete an order")
  w.WriteHeader(http.StatusNoContent)
}

