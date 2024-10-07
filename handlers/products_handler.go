package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"products_server/models"
	"products_server/repository"
	"strconv"
)

type ProductHandler struct {
	repo *repository.ProductRepository
}

func NewProductHandler(repo *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

func (h *ProductHandler) GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid product ID format", http.StatusBadRequest)
		return
	}

	product, err := h.repo.GetProductByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to get product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if product == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(product); err != nil {
		http.Error(w, "Failed to encode product: "+err.Error(), http.StatusInternalServerError)
	}
}

func (h *ProductHandler) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	page := 1
	pageSize := 10
	textFilter := r.URL.Query().Get("search")

	if r.URL.Query().Get("page") != "" {
		p, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err == nil && p > 0 {
			page = p
		}
	}

	if r.URL.Query().Get("page_size") != "" {
		ps, err := strconv.Atoi(r.URL.Query().Get("page_size"))
		if err == nil && ps > 0 {
			pageSize = ps
		}
	}

	products, totalRecords, err := h.repo.GetProducts(r.Context(), page, pageSize, textFilter)
	if err != nil {
		http.Error(w, "Failed to get products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	totalPages := (totalRecords + pageSize - 1) / pageSize

	response := struct {
		Page       int                     `json:"page"`
		PageSize   int                     `json:"page_size"`
		TotalPages int                     `json:"total_pages"`
		Data       []models.ProductDetails `json:"data"`
	}{
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
		Data:       products,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode products: "+err.Error(), http.StatusInternalServerError)
	}
}
