package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ternyavsky/tms_backend/internal/dto"
	"github.com/ternyavsky/tms_backend/internal/services"
)

type LinkController struct {
	linkService services.LinkService
}

func NewLinkController(linkService services.LinkService) *LinkController {
	return &LinkController{linkService: linkService}
}

// @Tags links
// @Produce  json
// @Success 200 {object} []domain.Link
// @Router /api/v1/links [get]
func (link *LinkController) GetAllLinks(ctx *gin.Context) {
	values, err := link.linkService.GetAllLinks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}

	ctx.JSON(200, gin.H{
		"data": values,
	})
}

// @Tags links
// @Produce  json
// @Param session body dto.LinkRequestDto true "Данные ссылки"
// @Success 201 {object} domain.Link
// @Router /api/v1/links [post]
func (link *LinkController) CreateLink(ctx *gin.Context) {
	var createLinkRequestDto dto.LinkRequestDto
	if err := ctx.ShouldBindJSON(&createLinkRequestDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	value, err := link.linkService.CreateLink(createLinkRequestDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}
	ctx.JSON(http.StatusCreated, value)
}
