package dto

import (
	"time"

	"github.com/ternyavsky/tms_backend/internal/domain"
)

type SessionRequestDto struct {
	ChannelId string `json:"channel_id"`
	SessionId string `json:"session_id"`
	SourceId  uint   `json:"source_id"`
	IsActive  bool   `json:"is_active"`
}
type SessionResponseDto struct {
	Id        uint      `json:"id"`
	ChannelId string    `json:"channel_id"`
	SessionId string    `json:"session_id"`
	IsActive  bool      `json:"is_active"`
	JoinAt    time.Time `json:"join_at"`
	CreatedAt time.Time `json:"created_at"`
}

func GetAllSessionsResponseDto(sessions []domain.Session) []SessionResponseDto {
	var values []SessionResponseDto
	for _, session := range sessions {
		values = append(values, SessionResponseDto{
			Id:        session.Id,
			ChannelId: session.ChannelId,
			SessionId: session.SessionId,
			IsActive:  session.IsActive,
			JoinAt:    session.JoinAt,
			CreatedAt: session.CreatedAt,
		})
	}
	return values
}

func CreateSessionResponseDto(session domain.Session) SessionResponseDto {
	return SessionResponseDto{
		Id:        session.Id,
		ChannelId: session.ChannelId,
		SessionId: session.SessionId,
		IsActive:  session.IsActive,
		JoinAt:    session.JoinAt,
		CreatedAt: session.CreatedAt,
	}
}
