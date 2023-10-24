package main

import (
	"log"
	"net/http"

	config "github.com/adhimaswaskita/go-product-management/configs"
	"github.com/adhimaswaskita/go-product-management/internal/database"
	"github.com/adhimaswaskita/go-product-management/internal/handlers"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

const (
	// ConfigFileLocation is the file configuration of ths service.
	ConfigFileLocation = "configs/config.yaml"
)

func SetupRouter(h handlers.IHandler) *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/admin", h.GetAdmin).Methods("GET")
	api.HandleFunc("/admin", h.CreateAdmin).Methods("POST")
	api.HandleFunc("/admin", h.UpdateAdmin).Methods("PUT")
	api.HandleFunc("/admin", h.DeleteAdmin).Methods("DELETE")

	api.HandleFunc("/product-category", h.GetProductCategory).Methods("GET")
	api.HandleFunc("/product-category", h.CreateProductCategory).Methods("POST")
	api.HandleFunc("/product-category", h.UpdateProductCategory).Methods("PUT")
	api.HandleFunc("/product-category", h.DeleteProductCategory).Methods("DELETE")

	api.HandleFunc("/product", h.GetProduct).Methods("GET")
	api.HandleFunc("/product", h.CreateProduct).Methods("POST")
	api.HandleFunc("/product", h.UpdateProduct).Methods("PUT")
	api.HandleFunc("/product", h.DeleteProduct).Methods("DELETE")

	return r
}

func main() {
	configLoader := config.NewYamlConfigLoader(ConfigFileLocation)
	conf, err := configLoader.GetServiceConfig()
	if err != nil {
		log.Fatalf("Unable to load configuration: %v", err)
	}

	log.Print("Success read config file")

	db, err := database.NewDB(&conf.SourceData)
	if err != nil {
		log.Fatalf("Unable to connect database : %v", err)
	}

	h := handlers.NewHandler(db, conf, validator.New())

	r := SetupRouter(h)

	log.Printf("Starting http server at %v", conf.ServiceData.Address)

	err = http.ListenAndServe(conf.ServiceData.Address, r)
	if err != nil {
		log.Fatalf("Unable to run http server: %v", err)
	}

}
