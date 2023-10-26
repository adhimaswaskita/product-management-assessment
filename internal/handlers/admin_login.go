package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/adhimaswaskita/go-product-management/internal/pkg/jwt"
	responsewriter "github.com/adhimaswaskita/go-product-management/internal/pkg/responseformatter"
)

type loginParam struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	rf := &responsewriter.ResponseFormat{}

	decoder := json.NewDecoder(r.Body)
	param := &loginParam{}

	err := decoder.Decode(param)
	if err != nil {
		log.Fatal(err)
		rf.ResponseNOK(http.StatusBadRequest, err, w)
		return
	}

	err = h.v.Struct(*param)
	if err != nil {
		log.Fatal(err)
		rf.ResponseNOK(http.StatusBadRequest, err, w)
		return
	}

	_, err = h.db.GetAdminByEmailAndPassword(param.Email, param.Password)
	if err != nil {
		log.Fatal(err)
		rf.ResponseNOK(http.StatusInternalServerError, err, w)
		return
	}

	token, err := jwt.GenerateNewTokenString()
	if err != nil {
		log.Fatal(err)
		rf.ResponseNOK(http.StatusInternalServerError, err, w)
		return
	}

	response := loginResponse{
		Token: token,
	}

	rf.ResponseOK(http.StatusOK, response, w)
}
