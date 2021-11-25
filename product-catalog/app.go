package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var (
	router *mux.Router
)

func InitRestServer(user, password, host, dbname, addr string) {
	connectDb(user, password, host, dbname)

	router = mux.NewRouter()
	initializeRoutes()

	log.Fatal(http.ListenAndServe(addr, router))
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s", "getProducts")

	products := Products{}
	err := products.getProducts()
	if err != nil {
		log.Fatal(err)
	}

	respondWithJSON(w, http.StatusOK, products)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s", "createProduct")

	var p Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := p.createProduct(); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s", "updateProduct")

	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var p Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	p.ID = id.String()

	if err := p.updateProduct(); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s", "deleteProduct")

	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var p Product
	p.ID = id.String()

	if err := p.deleteProduct(); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func initializeRoutes() {
	router.HandleFunc("/v1/products", getProducts).Methods("GET")
	router.HandleFunc("/v1/products", createProduct).Methods("POST")
	router.HandleFunc("/v1/products/{id}", updateProduct).Methods("PUT")
	router.HandleFunc("/v1/products/{id}", deleteProduct).Methods("DELETE")
}
