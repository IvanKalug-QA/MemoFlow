package memo

import (
	"fmt"
	"memoflow/configs"
	"net/http"
)

type MemoHandlerDeps struct {
	*configs.Config
}

type MemoHandler struct {
	*configs.Config
}

func (m *MemoHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Create")
	}
}

func (m *MemoHandler) Read() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Read")
	}
}

func (m *MemoHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Update")
	}
}

func (m *MemoHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Delete")
	}
}

func NewMemoHandler(router *http.ServeMux, deps MemoHandlerDeps) {
	handler := &MemoHandler{Config: deps.Config}
	router.HandleFunc("POST /memo", handler.Create())
	router.HandleFunc("GET /memo", handler.Read())
	router.HandleFunc("DELETE /memo/{id}", handler.Delete())
	router.HandleFunc("PATCH /memo/{id}", handler.Update())
}
