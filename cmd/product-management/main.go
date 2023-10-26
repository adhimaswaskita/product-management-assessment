package main

import (
	"log"
	"net/http"

	config "github.com/adhimaswaskita/go-product-management/configs"
	"github.com/adhimaswaskita/go-product-management/internal/database"
	"github.com/adhimaswaskita/go-product-management/internal/handlers"
	m "github.com/adhimaswaskita/go-product-management/internal/middlewares"
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

	api.HandleFunc("/login", h.Login).Methods(("POST"))
	api.Handle("/admin", m.AuthMiddleware(http.HandlerFunc(h.GetAdmin))).Methods("GET")
	api.Handle("/admin", m.AuthMiddleware(http.HandlerFunc(h.CreateAdmin))).Methods("POST")
	api.Handle("/admin", m.AuthMiddleware(http.HandlerFunc(h.UpdateAdmin))).Methods("PUT")
	api.Handle("/admin", m.AuthMiddleware(http.HandlerFunc(h.DeleteAdmin))).Methods("DELETE")

	api.Handle("/product-category", m.AuthMiddleware(http.HandlerFunc(h.GetProductCategory))).Methods("GET")
	api.Handle("/product-category", m.AuthMiddleware(http.HandlerFunc(h.CreateProductCategory))).Methods("POST")
	api.Handle("/product-category", m.AuthMiddleware(http.HandlerFunc(h.UpdateProductCategory))).Methods("PUT")
	api.Handle("/product-category", m.AuthMiddleware(http.HandlerFunc(h.DeleteProductCategory))).Methods("DELETE")

	api.Handle("/product", m.AuthMiddleware(http.HandlerFunc(h.GetProduct))).Methods("GET")
	api.Handle("/product", m.AuthMiddleware(http.HandlerFunc(h.CreateProduct))).Methods("POST")
	api.Handle("/product", m.AuthMiddleware(http.HandlerFunc(h.UpdateProduct))).Methods("PUT")
	api.Handle("/product", m.AuthMiddleware(http.HandlerFunc(h.DeleteProduct))).Methods("DELETE")

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
