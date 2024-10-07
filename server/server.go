package server

import (
	gorillahandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"net/http"
	"products_server/handlers"
	"products_server/repository"
)

func NewServer(dbPool *pgxpool.Pool) *http.Server {
	productRepo := repository.NewProductRepository(dbPool)
	productHandler := handlers.NewProductHandler(productRepo)

	r := mux.NewRouter()

	apiRouter := r.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/products", productHandler.GetProductsHandler).Methods("GET")
	apiRouter.HandleFunc("/products/{id}", productHandler.GetProductByIDHandler).Methods("GET")

	apiRouter.HandleFunc("/health", handlers.HealthHandler).Methods("GET")

	corsOptions := gorillahandlers.AllowedOrigins([]string{"*"})
	corsMethods := gorillahandlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE"})

	return &http.Server{
		Addr:    ":8080",
		Handler: gorillahandlers.CORS(corsOptions, corsMethods)(r),
	}
}
