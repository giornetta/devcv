package server

import (
	"encoding/json"
	"net/http"

	"github.com/giornetta/devcv/auth"
	"github.com/giornetta/devcv/developers"

	"github.com/go-chi/chi"
)

type developersHandler struct {
	s developers.Service
	a auth.Service
}

func (h *developersHandler) routes() *chi.Mux {
	mux := chi.NewMux()

	mux.Post("/login", h.login)

	mux.Post("/", h.create)

	mux.Get("/{username}", h.get)

	mux.Group(func(r chi.Router) {
		r.Use(authMiddleware(h.a))

		r.Put("/{username}", h.update)
		r.Delete("/{username}", h.delete)
	})

	return mux
}

func (h *developersHandler) login(w http.ResponseWriter, r *http.Request) {
	var req developers.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond(w, http.StatusBadRequest, e(err.Error()))
		return
	}

	res, err := h.s.Login(&req)
	if err != nil {
		respond(w, http.StatusInternalServerError, e(err.Error()))
		return
	}

	respond(w, http.StatusOK, res)
}

func (h *developersHandler) create(w http.ResponseWriter, r *http.Request) {
	var req developers.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond(w, http.StatusBadRequest, e(err.Error()))
		return
	}

	res, err := h.s.Create(&req)
	if err != nil {
		respond(w, http.StatusBadRequest, e(err.Error()))
		return
	}

	respond(w, http.StatusOK, res)
}

func (h *developersHandler) get(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	res, err := h.s.Get(&developers.UsernameRequest{Username: username})
	if err != nil {
		respond(w, http.StatusBadRequest, e(err.Error()))
		return
	}

	respond(w, http.StatusOK, res)
}

func (h *developersHandler) update(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	var req developers.Developer
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond(w, http.StatusBadRequest, e(err.Error()))
		return
	}

	req.Username = username

	err := h.s.Update(req)
	if err != nil {
		respond(w, http.StatusBadRequest, e(err.Error()))
		return
	}

	respond(w, http.StatusOK, nil)
}

func (h *developersHandler) delete(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	err := h.s.Delete(&developers.UsernameRequest{Username: username})
	if err != nil {
		respond(w, http.StatusBadRequest, e(err.Error()))
		return
	}

	respond(w, http.StatusOK, nil)
}
