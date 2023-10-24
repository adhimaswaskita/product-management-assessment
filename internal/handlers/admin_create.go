package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/adhimaswaskita/go-product-management/internal/models"
	responsewriter "github.com/adhimaswaskita/go-product-management/internal/pkg/responseformatter"
)

func (h *Handler) CreateAdmin(w http.ResponseWriter, r *http.Request) {
	rf := &responsewriter.ResponseFormat{}

	decoder := json.NewDecoder(r.Body)
	param := &models.Admin{}

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

	err = h.db.CreateAdmin(*param)
	if err != nil {
		log.Fatal(err)
		rf.ResponseNOK(http.StatusInternalServerError, err, w)
		return
	}

	log.Print("Success create admin")

	rf.ResponseOK(http.StatusOK, "Success add admin", w)
}
