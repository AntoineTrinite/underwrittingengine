package handler

import (
	"net/http"
	"uwe/types"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type User struct {
	Name string
}

func HandleGetCustomer(w http.ResponseWriter, r *http.Request) error {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		return NewAPIError(http.StatusBadRequest, err)
	}
	customer := types.Customer{
		ID: id,
	}
	return writeJSON(w, http.StatusOK, customer)
}
