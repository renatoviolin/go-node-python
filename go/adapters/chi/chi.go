package chi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/renatoviolin/go-node/dto"
	"github.com/renatoviolin/go-node/usecase"
)

type ChiHandler struct {
	chi *chi.Mux
	uc  *usecase.FindAll
}

func NewHandler(uc *usecase.FindAll) *ChiHandler {
	chiRouter := chi.NewRouter()
	return &ChiHandler{
		chi: chiRouter,
		uc:  uc,
	}
}

func (h *ChiHandler) SetupAndRun(address string) {
	h.chi.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	h.chi.Get("/single_json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dto.NewPayload())
	})

	h.chi.Get("/mongo_json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		result, err := h.uc.Execute()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(result)
	})

	fmt.Printf("Chi running at %s\n", address)
	http.ListenAndServe(address, h.chi)
}
