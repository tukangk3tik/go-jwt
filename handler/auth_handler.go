package handler

import (
	"encoding/json"
	"go-jwt/dto"
	"go-jwt/helper"
	"go-jwt/service"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginData dto.LoginDto

	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		helper.ReturnFailResponse(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if loginData.Email != "admin@mail.com" || loginData.Password != "admin" {
		helper.ReturnFailResponse(w, "Invalid credentials", http.StatusNotFound)
		return
	}

	token := service.JwtServiceInstance.GenerateToken(loginData.Email)
	helper.ReturnSuccessResponse[dto.TokenDto](w, token, http.StatusCreated)
}
