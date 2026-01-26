package memo

import (
	"fmt"
	"memoflow/pkg/req"
	"memoflow/pkg/res"
	"net/http"
)

type MemoHandlerDeps struct {
	MemoResository *MemoRepository
}

type MemoHandler struct {
	MemoResository *MemoRepository
}

func (m *MemoHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[MemoRequest](&w, r)
		if err != nil {
			return
		}
		memo := NewMemo(body)
		createdMemo, err := m.MemoResository.Create(memo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		fmt.Println(createdMemo)
		res.Json(w, createdMemo, http.StatusCreated)
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
