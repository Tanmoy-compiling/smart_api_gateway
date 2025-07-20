package product

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

// Product represents a product entity
type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

// Service interface defines the product service operations
type Service interface {
	CreateProduct(ctx context.Context, product *Product) error
	GetProduct(ctx context.Context, id string) (*Product, error)
	UpdateProduct(ctx context.Context, product *Product) error
	DeleteProduct(ctx context.Context, id string) error
	ListProducts(ctx context.Context) ([]*Product, error)
}

// Repository interface defines the product storage operations
type Repository interface {
	Save(ctx context.Context, product *Product) error
	FindByID(ctx context.Context, id string) (*Product, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id string) error
	FindAll(ctx context.Context) ([]*Product, error)
}

type service struct {
	repo Repository
}

// NewService creates a new product service instance
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateProduct(ctx context.Context, product *Product) error {
	if product.Name == "" {
		return errors.New("product name is required")
	}
	if product.Price < 0 {
		return errors.New("product price must be non-negative")
	}
	if product.Stock < 0 {
		return errors.New("product stock must be non-negative")
	}

	product.ID = uuid.New().String()
	return s.repo.Save(ctx, product)
}

func (s *service) GetProduct(ctx context.Context, id string) (*Product, error) {
	if id == "" {
		return nil, errors.New("product id is required")
	}
	return s.repo.FindByID(ctx, id)
}

func (s *service) UpdateProduct(ctx context.Context, product *Product) error {
	if product.ID == "" {
		return errors.New("product id is required")
	}
	if product.Name == "" {
		return errors.New("product name is required")
	}
	if product.Price < 0 {
		return errors.New("product price must be non-negative")
	}
	if product.Stock < 0 {
		return errors.New("product stock must be non-negative")
	}

	return s.repo.Update(ctx, product)
}

func (s *service) DeleteProduct(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("product id is required")
	}
	return s.repo.Delete(ctx, id)
}

func (s *service) ListProducts(ctx context.Context) ([]*Product, error) {
	return s.repo.FindAll(ctx)
}
