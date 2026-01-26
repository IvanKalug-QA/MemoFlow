package memo

import (
	"fmt"
	"net/http"
)

type MemoHandlerDeps struct {
	MemoResository *MemoResository
}

type MemoHandler struct {
	MemoResository *MemoResository
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
	handler := &MemoHandler{MemoResository: deps.MemoResository}
	router.HandleFunc("POST /memo", handler.Create())
	router.HandleFunc("GET /memo", handler.Read())
	router.HandleFunc("DELETE /memo/{id}", handler.Delete())
	router.HandleFunc("PATCH /memo/{id}", handler.Update())
}
