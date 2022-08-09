package handler

import (
	"go-jwt/helper"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	helper.ReturnSuccessResponse(w, "Welcome to home!", 200)
}
