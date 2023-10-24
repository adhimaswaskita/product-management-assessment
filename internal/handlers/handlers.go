package handlers

import (
	"net/http"

	config "github.com/adhimaswaskita/go-product-management/configs"
	"github.com/adhimaswaskita/go-product-management/internal/database"
	validator "github.com/go-playground/validator"
)

// Handler hold the function handler for API's endpoint.
type IHandler interface {
	CreateAdmin(w http.ResponseWriter, r *http.Request)
	GetAdmin(w http.ResponseWriter, r *http.Request)
	UpdateAdmin(w http.ResponseWriter, r *http.Request)
	DeleteAdmin(w http.ResponseWriter, r *http.Request)

	CreateProductCategory(w http.ResponseWriter, r *http.Request)
	GetProductCategory(w http.ResponseWriter, r *http.Request)
	UpdateProductCategory(w http.ResponseWriter, r *http.Request)
	DeleteProductCategory(w http.ResponseWriter, r *http.Request)

	CreateProduct(w http.ResponseWriter, r *http.Request)
	GetProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
}

// Handler holds the API endpoint's function handler.
type Handler struct {
	db     database.IDB
	config *config.ServiceConfig
	v      *validator.Validate
}

// NewHandler function to make connection database into handler
func NewHandler(db database.IDB, config *config.ServiceConfig, validator *validator.Validate) *Handler {
	return &Handler{
		db:     db,
		config: config,
		v:      validator,
	}
}
