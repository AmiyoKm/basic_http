package user

import (
	"fmt"
	"net/http"

	"github.com/AmiyoKm/basic_http/config"
	"github.com/AmiyoKm/basic_http/domain"
	"github.com/AmiyoKm/basic_http/jwt"
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

type registerUserPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var payload registerUserPayload
	utils.ReadJSON(r, &payload)

	pass := domain.Password{
		String: payload.Password,
	}
	pass.Hash()

	user := &domain.Users{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: pass,
	}

	user, err := h.svc.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error creating user , err :%s", err.Error())
		return
	}
	env := utils.Envelop{Message: "user created successfully", Value: user}

	utils.WriteJSON(w, env)
}

type loginUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var payload loginUserPayload
	utils.ReadJSON(r, &payload)

	pass := domain.Password{
		String: payload.Password,
	}
	pass.Hash()

	user, err := h.svc.GetByEmail(payload.Email)
	if err != nil {
		fmt.Fprintf(w, "error user not found, err :%s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !user.Password.Match(payload.Password) {
		fmt.Fprintf(w, "error password does not match")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newJwt, err := jwt.NewJWT(user.ID, h.cnf.JWTSecretKey)
	if err != nil {

		fmt.Fprintf(w, "error creating jwt")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token, err := newJwt.ToString()
	if err != nil {
		fmt.Fprintf(w, "error creating jwt to string")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	env := utils.Envelop{Message: "user logged in successfully", Value: token}

	utils.WriteJSON(w, env)
}
