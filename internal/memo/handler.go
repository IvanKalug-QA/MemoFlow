package memo

import (
	"memoflow/pkg/middleware"
	"memoflow/pkg/req"
	"memoflow/pkg/res"
	"net/http"
	"strconv"

	"gorm.io/gorm"
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
		res.Json(w, createdMemo, http.StatusCreated)
	}
}

func (m *MemoHandler) Read() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		memo, err := m.MemoResository.GetByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		res.Json(w, memo, http.StatusOK)
	}
}

func (m *MemoHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[MemoUpdateRequest](&w, r)
		if err != nil {
			return
		}
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		memo, err := m.MemoResository.Update(&Memo{
			Model:       gorm.Model{ID: uint(id)},
			Name:        body.Name,
			Description: body.Description,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, memo, http.StatusOK)
	}
}

func (m *MemoHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if _, err = m.MemoResository.GetByID(id); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = m.MemoResository.Delete(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w, nil, 200)
	}
}

func NewMemoHandler(router *http.ServeMux, deps MemoHandlerDeps) {
	handler := &MemoHandler{MemoResository: deps.MemoResository}
	router.HandleFunc("POST /memo", handler.Create())
	router.HandleFunc("GET /memo/{id}", handler.Read())
	router.HandleFunc("DELETE /memo/{id}", handler.Delete())
	router.Handle("PATCH /memo/{id}", middleware.IsAuthed(handler.Update()))
}
