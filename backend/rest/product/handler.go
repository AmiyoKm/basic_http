package product

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/AmiyoKm/basic_http/config"
	"github.com/AmiyoKm/basic_http/domain"
	"github.com/AmiyoKm/basic_http/utils"
)

type Handler struct {
	cnf *config.Config
	svc Service
}

func NewHandler(cnf *config.Config, svc Service) *Handler {
	return &Handler{
		cnf: cnf,
		svc: svc,
	}
}

func (h *Handler) getProductByID(w http.ResponseWriter, r *http.Request) {

	paramID := r.PathValue("id")
	product, err := h.svc.GetByID(paramID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, utils.Envelop{Message: "product created", Value: product})
}

type pagination struct {
	Data     []*domain.Product `json:"data"`
	Metadata Metadata          `json:"metadata"`
}

type Metadata struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}

func (h *Handler) getProducts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	pageStr := query.Get("page")
	limitStr := query.Get("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || pageStr == "" {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limitStr == "" {
		limit = 10
	}

	products, err := h.svc.Get(page, limit)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cnt, err := h.svc.Count()

	pagination := pagination{
		Data: products,
		Metadata: Metadata{
			Page:       page,
			Limit:      limit,
			TotalItems: cnt,
			TotalPages: (cnt / limit) + 1,
		},
	}

	utils.WriteJSON(w, utils.Envelop{Message: "success", Value: pagination})
}

func (h *Handler) createProduct(w http.ResponseWriter, r *http.Request) {
	var product domain.Product

	utils.ReadJSON(r, &product)

	err := h.svc.Create(&product)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, utils.Envelop{Message: "product created", Value: product})
}

func (h *Handler) updateProduct(w http.ResponseWriter, r *http.Request) {
	var product domain.Product
	utils.ReadJSON(r, &product)

	paramID := r.PathValue("id")
	product.ID = paramID

	err := h.svc.Update(&product)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, utils.Envelop{Message: "product updated", Value: product})
}

func (h *Handler) deleteProduct(w http.ResponseWriter, r *http.Request) {
	paramID := r.PathValue("id")

	err := h.svc.Delete(paramID)
	if err != nil {
		log.Println(err)
		utils.WriteJSON(w, utils.Envelop{Message: "product not found"})
		return
	}
	utils.WriteJSON(w, utils.Envelop{Message: "product deleted"})
}
