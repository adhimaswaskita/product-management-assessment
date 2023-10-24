package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/adhimaswaskita/go-product-management/internal/models"
	responsewriter "github.com/adhimaswaskita/go-product-management/internal/pkg/responseformatter"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	rf := &responsewriter.ResponseFormat{}

	decoder := json.NewDecoder(r.Body)
	param := &models.Product{}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		rf.ResponseNOK(http.StatusBadRequest, err, w)
		return
	}

	err = decoder.Decode(param)
	if err != nil {
		log.Fatal(err)
		rf.ResponseNOK(http.StatusBadRequest, err, w)
		return
	}

	err = h.db.UpdateProduct(uint(id), *param)
	if err != nil {
		log.Fatal(err)
		rf.ResponseNOK(http.StatusInternalServerError, err, w)
		return
	}

	log.Print("Success update product")

	rf.ResponseOK(http.StatusOK, "Success update product", w)
}
