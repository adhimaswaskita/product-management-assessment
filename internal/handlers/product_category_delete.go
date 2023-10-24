package handlers

import (
	"log"
	"net/http"
	"strconv"

	responsewriter "github.com/adhimaswaskita/go-product-management/internal/pkg/responseformatter"
)

func (h *Handler) DeleteProductCategory(w http.ResponseWriter, r *http.Request) {
	rf := &responsewriter.ResponseFormat{}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		rf.ResponseNOK(http.StatusBadRequest, err, w)
		return
	}

	err = h.db.DeleteProductCategory(uint(id))
	if err != nil {
		log.Fatal(err)
		rf.ResponseNOK(http.StatusInternalServerError, err, w)
		return
	}

	log.Print("Success delete product category")

	rf.ResponseOK(http.StatusOK, "Success delete product category", w)
}
