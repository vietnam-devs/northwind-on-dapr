package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	Db     *sqlx.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connString :=
		fmt.Sprintf("user=%s password=%s host=localhost dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.Db, err = sqlx.Connect("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	initDb(a.Db)

	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s", "getProducts")

	products := Products{}
	err := products.getProducts(a.Db)
	if err != nil {
		log.Fatal(err)
	}

	respondWithJSON(w, http.StatusOK, products)
}

func (a *App) createProduct(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s", "createProduct")

	var p Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := p.createProduct(a.Db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}

func (a *App) updateProduct(w http.ResponseWriter, r *http.Request) {
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

	if err := p.updateProduct(a.Db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func (a *App) deleteProduct(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s", "deleteProduct")

	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var p Product
	p.ID = id.String()

	if err := p.deleteProduct(a.Db); err != nil {
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

func (a *App) RunApp(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/v1/products", a.getProducts).Methods("GET")
	a.Router.HandleFunc("/v1/products", a.createProduct).Methods("POST")
	a.Router.HandleFunc("/v1/products/{id}", a.updateProduct).Methods("PUT")
	a.Router.HandleFunc("/v1/products/{id}", a.deleteProduct).Methods("DELETE")
}
