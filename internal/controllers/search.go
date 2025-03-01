package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ternyavsky/tms_backend/internal/services"
)

type SearchController struct {
	searchService services.SearchService
}

func NewSearchController(searchService services.SearchService) *SearchController {
	return &SearchController{searchService: searchService}
}

// @Tags search
// @Produce  json
// @Param channelUrl path string true "URL канала"
// @Success 200 {object} domain.Link
// Failure 400 {object} map[string]string
// Failure 404 {object} map[string]string
// @Router /api/v1/search/{channelUrl} [get]
func (search *SearchController) SearchChannel(ctx *gin.Context) {
	channelUrl := ctx.Param("channelUrl")
	link, err := search.searchService.SearchChannel(channelUrl)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": link,
	})
}

// @Tags search
// @Produce  json
// @Param linkId path number true "ID ссылки"
// @Success 200 {object} domain.Link
// Failure 400 {object} map[string]string
// Failure 404 {object} map[string]string
// @Router /api/v1/search/status/{linkId} [get]
func (search *SearchController) GetSearchStatus(ctx *gin.Context) {
	idStr := ctx.Param("linkId")
	linkId, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	link, err := search.searchService.GetSearchStatus(uint(linkId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": link,
	})
}
