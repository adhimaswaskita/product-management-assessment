package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/adhimaswaskita/go-product-management/internal/models"
	responsewriter "github.com/adhimaswaskita/go-product-management/internal/pkg/responseformatter"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	rf := &responsewriter.ResponseFormat{}

	decoder := json.NewDecoder(r.Body)
	param := &models.Product{}

	err := decoder.Decode(param)
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

	err = h.db.CreateProduct(*param)
	if err != nil {
		log.Fatal(err)
		rf.ResponseNOK(http.StatusInternalServerError, err, w)
		return
	}

	log.Print("Success create product")

	rf.ResponseOK(http.StatusOK, "Success add product", w)
}
