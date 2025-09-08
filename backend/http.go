package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type envelop struct {
	Message string
	Value   any
}

func createUser(w http.ResponseWriter, r *http.Request) {
	type Payload struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var userPayload Payload
	readJSON(r, &userPayload)

	pass := Password{
		String: userPayload.Password,
	}
	pass.Hash()

	user := User{
		Email:    userPayload.Email,
		Name:     userPayload.Name,
		Password: pass,
	}

	userStorage = append(userStorage, user)

	jwt, err := NewJWT(fmt.Sprint(len(userStorage)), SECRET)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error creating jwt , err :%s", err.Error())
		return
	}

	jwtHash, err := jwt.ToString()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error creating jwt , err :%s", err.Error())
		return
	}

	env := envelop{
		Message: "successfully created user",
		Value:   jwtHash,
	}

	writeJSON(w, env)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := productStorage.Get()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, products)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var product Product

	readJSON(r, &product)

	err := productStorage.Create(&product)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, envelop{Message: "product created", Value: product})

}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	readJSON(r, &product)

	paramID := r.PathValue("id")
	err := productStorage.Update(paramID, &product)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, envelop{Message: "product updated", Value: product})
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	paramID := r.PathValue("id")

	err := productStorage.Delete(paramID)
	if err != nil {
		log.Println(err)
		writeJSON(w, envelop{Message: "product not found"})
		return
	}
	writeJSON(w, envelop{Message: "product deleted"})
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	type HealthStatus struct {
		Status    string    `json:"status"`
		Timestamp time.Time `json:"timestamp"`
		Version   string    `json:"version"`
	}

	writeJSON(w, HealthStatus{
		Status:    "healthy",
		Timestamp: time.Now(),
		Version:   "1.0.0",
	})
}
