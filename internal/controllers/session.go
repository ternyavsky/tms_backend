package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ternyavsky/tms_backend/internal/dto"
	"github.com/ternyavsky/tms_backend/internal/services"
)

type SessionConrtoller struct {
	sessionService services.SessionService
}

func NewSessionController(sessionService services.SessionService) *SessionConrtoller {
	return &SessionConrtoller{sessionService: sessionService}
}

// @Tags sessions
// @Produce  json
// @Success 200 {object} []domain.Session
// @Router /api/v1/sessions [get]
func (session *SessionConrtoller) GetAllSessions(ctx *gin.Context) {
	values, err := session.sessionService.GetAllSessions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": values,
	})
}

// @Tags sessions
// @Produce  json
// @Param session body dto.SessionRequestDto true "Данные сессии"
// @Success 201 {object} domain.Session
// @Router /api/v1/sessions [post]
func (session *SessionConrtoller) CreateSession(ctx *gin.Context) {
	var createSessionRequestDto dto.SessionRequestDto
	if err := ctx.ShouldBindJSON(&createSessionRequestDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	value, err := session.sessionService.CreateSession(createSessionRequestDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, value)
}
