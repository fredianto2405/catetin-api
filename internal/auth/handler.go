package auth

import (
	"github.com/fredianto2405/catetin-api/pkg/errors"
	"github.com/fredianto2405/catetin-api/pkg/jwt"
	"github.com/fredianto2405/catetin-api/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Respond(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	if err := errors.Validate.Struct(request); err != nil {
		response.Respond(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	user, err := h.service.Login(&request)
	if err != nil {
		response.Respond(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	token, err := jwt.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		response.Respond(c, http.StatusInternalServerError, false, "failed to generate token", nil, nil)
		return
	}

	data := LoginResponse{
		Email:       user.Email,
		AccessToken: token,
		Role:        user.Role,
	}

	response.Respond(c, http.StatusOK, true, "login berhasil", data, nil)
}

func (h *Handler) ChangePassword(c *gin.Context) {
	claims, ok := jwt.GetUserClaims(c)
	if !ok {
		response.Respond(c, http.StatusUnauthorized, false, "user not found in ctx", nil, nil)
		return
	}
	email := claims.Email

	var request ChangePasswordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Respond(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	if err := errors.Validate.Struct(request); err != nil {
		response.Respond(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	err := h.service.ChangePassword(email, &request)
	if err != nil {
		response.Respond(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	response.Respond(c, http.StatusOK, true, "password berhasil diubah", nil, nil)
}
