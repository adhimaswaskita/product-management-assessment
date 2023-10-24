package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/adhimaswaskita/go-product-management/internal/models"
	responsewriter "github.com/adhimaswaskita/go-product-management/internal/pkg/responseformatter"
)

func (h *Handler) UpdateAdmin(w http.ResponseWriter, r *http.Request) {
	rf := &responsewriter.ResponseFormat{}

	decoder := json.NewDecoder(r.Body)
	param := &models.Admin{}

	adminID, err := strconv.Atoi(r.FormValue("id"))
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

	err = h.v.Struct(param)
	if err != nil {
		log.Fatal(err)
		rf.ResponseNOK(http.StatusBadRequest, err, w)
		return
	}

	err = h.db.UpdateAdmin(uint(adminID), *param)
	if err != nil {
		log.Fatal(err)
		rf.ResponseNOK(http.StatusInternalServerError, err, w)
		return
	}

	log.Print("Success update admin")

	rf.ResponseOK(http.StatusOK, "Success update admin", w)
}
