package database

import (
	config "github.com/adhimaswaskita/go-product-management/configs"
	mysql "github.com/adhimaswaskita/go-product-management/internal/database/mySQL"
	"github.com/adhimaswaskita/go-product-management/internal/models"
)

type IDB interface {
	CreateAdmin(models.Admin) error
	GetAdmin() ([]models.Admin, error)
	GetAdminByEmailAndPassword(email, password string) (*models.Admin, error)
	UpdateAdmin(uint, models.Admin) error
	DeleteAdmin(id uint) error

	CreateProductCategory(models.ProductCategory) error
	GetProductCategory() ([]models.ProductCategory, error)
	UpdateProductCategory(uint, models.ProductCategory) error
	DeleteProductCategory(id uint) error

	CreateProduct(models.Product) error
	GetProduct() ([]models.Product, error)
	UpdateProduct(uint, models.Product) error
	DeleteProduct(id uint) error
}

// NewDB initialize selected DB for bec-user
func NewDB(config *config.SourceDataConfig) (IDB, error) {
	db, err := mysql.NewMySQL(config.DBUsername, config.DBPassword, config.DBName, config.DBServer, config.DBPort)
	if err != nil {
		return nil, err
	}

	db.DB.AutoMigrate(&models.Admin{}, &models.Product{}, &models.ProductCategory{}, &models.TransactionHistory{})

	return db, nil
}
