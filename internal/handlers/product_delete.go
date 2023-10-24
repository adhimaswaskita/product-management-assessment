package handlers

import (
	"log"
	"net/http"
	"strconv"

	responsewriter "github.com/adhimaswaskita/go-product-management/internal/pkg/responseformatter"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	rf := &responsewriter.ResponseFormat{}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		rf.ResponseNOK(http.StatusBadRequest, err, w)
		return
	}

	err = h.db.DeleteProduct(uint(id))
	if err != nil {
		log.Fatal(err)
		rf.ResponseNOK(http.StatusInternalServerError, err, w)
		return
	}

	log.Print("Success delete product")

	rf.ResponseOK(http.StatusOK, "Success delete product", w)
}
