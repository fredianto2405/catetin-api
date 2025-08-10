package member

import (
	"github.com/fredianto2405/catetin-api/pkg/errors"
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

func (h *Handler) CreateMember(c *gin.Context) {
	var request CreateMemberRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	if err := errors.Validate.Struct(request); err != nil {
		c.Error(err)
		return
	}

	if err := h.service.Create(&request); err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusCreated, true, "member berhasil ditambahkan", nil, nil)
}

func (h *Handler) GetMembers(c *gin.Context) {
	members, err := h.service.GetAll()
	if err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, "data member berhasil diambil", members, nil)
}

func (h *Handler) UpdateMember(c *gin.Context) {
	id := c.Param("id")
	var request UpdateMemberRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	if err := errors.Validate.Struct(&request); err != nil {
		c.Error(err)
		return
	}

	if err := h.service.Update(id, &request); err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, "data member berhasil diperbarui", nil, nil)
}

func (h *Handler) DeleteMember(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, "member berhasil dihapus", nil, nil)
}
