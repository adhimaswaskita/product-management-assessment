package handlers

import (
	"log"
	"net/http"
	"strconv"

	responsewriter "github.com/adhimaswaskita/go-product-management/internal/pkg/responseformatter"
)

func (h *Handler) DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	rf := &responsewriter.ResponseFormat{}

	adminID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		rf.ResponseNOK(http.StatusBadRequest, err, w)
		return
	}

	err = h.db.DeleteAdmin(uint(adminID))
	if err != nil {
		log.Fatal(err)
		rf.ResponseNOK(http.StatusInternalServerError, err, w)
		return
	}

	log.Print("Success delete admin")

	rf.ResponseOK(http.StatusOK, "Success delete admin", w)
}
