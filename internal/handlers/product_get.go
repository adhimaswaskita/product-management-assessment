package handlers

import (
	"log"
	"net/http"

	responsewriter "github.com/adhimaswaskita/go-product-management/internal/pkg/responseformatter"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	rf := &responsewriter.ResponseFormat{}

	data, err := h.db.GetProduct()
	if err != nil {
		rf.ResponseNOK(http.StatusInternalServerError, err, w)
	}

	log.Print("Success get product")

	rf.ResponseOK(http.StatusOK, data, w)
}
