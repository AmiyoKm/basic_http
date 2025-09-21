package product

import (
	"log"
	"net/http"

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

func (h *Handler) getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.svc.Get()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, utils.Envelop{Message: "success", Value: products})
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
