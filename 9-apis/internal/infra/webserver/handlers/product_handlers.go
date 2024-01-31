package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/jfromjefferson/gi-course-9/internal/dto"
	"github.com/jfromjefferson/gi-course-9/internal/entity"
	"github.com/jfromjefferson/gi-course-9/internal/infra/database"
	pkgentity "github.com/jfromjefferson/gi-course-9/pkg/entity"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

type Error struct {
	Message string `json:"message"`
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// CreateProduct Create product godoc
// @Summary Create product
// @Description Create product...
// @Tags products
// @Accept json
// @Produce json
// @Param request body dto.CreateProductInput true "product request"
// @Success 201
// @Failure 500 {object} Error
// @Router /products [post]
// @Security ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := entity.NewProduct(product.Name, product.Code, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetProduct Get product godoc
// @Summary Get product
// @Description Get product...
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "product id" Format(uuid)
// @Success 200 {object} entity.Product
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /products/{id} [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	productID, err := pkgentity.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	product, err := h.ProductDB.GetByID(productID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// GetProducts Get products godoc
// @Summary Get products
// @Description Get products...
// @Tags products
// @Accept json
// @Produce json
// @Param page query string false "page number"
// @Param limit query string false "limit"
// @Success 200 {array} entity.Product
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /products [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	if page == "" {
		page = "1"
	}

	if limit == "" {
		limit = "0"
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	products, err := h.ProductDB.FindAll(int(pageInt), int(limitInt), sort)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// UpdateProduct Update product godoc
// @Summary Update product
// @Description Update product...
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "product id" Format(uuid)
// @Param request body dto.UpdateProductInput true "product request"
// @Success 200
// @Failure 404
// @Failure 500 {object} Error
// @Router /products/{id} [put]
// @Security ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	productID, err := pkgentity.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var product dto.UpdateProductInput
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	productFound, err := h.ProductDB.GetByID(productID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	productFound.Name = product.Name
	productFound.Code = product.Code
	productFound.Price = product.Price

	err = productFound.ValidateProduct()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.ProductDB.Update(productID, productFound)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteProduct Delete product godoc
// @Summary Delete product
// @Description Delete product...
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "product id" Format(uuid)
// @Success 200
// @Failure 404
// @Failure 500 {object} Error
// @Router /products/{id} [delete]
// @Security ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	productID, err := pkgentity.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	product, err := h.ProductDB.GetByID(productID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Delete(productID, product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
