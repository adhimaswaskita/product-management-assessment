package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	responsewriter "github.com/adhimaswaskita/go-product-management/internal/pkg/responseformatter"
)

type updateProductStockParam struct {
	Stock int `json:"stock" validate:"required"`
}

func (h *Handler) UpdateProductStock(w http.ResponseWriter, r *http.Request) {
	rf := &responsewriter.ResponseFormat{}

	decoder := json.NewDecoder(r.Body)
	param := &updateProductStockParam{}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		rf.ResponseNOK(http.StatusBadRequest, err, w)
		return
	}

	activity := r.FormValue("type")

	err = decoder.Decode(param)
	if err != nil {
		log.Fatal(err)
		rf.ResponseNOK(http.StatusBadRequest, err, w)
		return
	}

	err = h.v.Struct(param)
	if err != nil {
		log.Fatal(err)
		rf.ResponseNOK(http.StatusBadRequest, err, w)
		return
	}

	if activity == "out" {
		param.Stock *= -1
	}

	err = h.db.UpdateProductStock(uint(id), param.Stock)
	if err != nil {
		log.Fatal(err)
		rf.ResponseNOK(http.StatusInternalServerError, err, w)
		return
	}

	log.Print("Success update product")

	rf.ResponseOK(http.StatusOK, "Success update product", w)
}
