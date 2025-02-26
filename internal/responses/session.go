package responses

import (
	"time"

	"github.com/ternyavsky/tms_backend/internal/domain"
)

type SessionResponse struct {
		Id uint `json:"id"`
		ChannelId string    `json:"channel_id"`
		SessionId string    `json:"session_id"`
		IsActive  bool      `json:"is_active"`
		JoinAt    time.Time `json:"join_at"`
		CreatedAt time.Time `json:"created_at"`
}

func GetAllSessionsResponse(sessions []domain.Session) []SessionResponse {
	var values []SessionResponse
	for _, session := range sessions{
		values = append(values, SessionResponse{
			Id: session.Id, 
			ChannelId: session.ChannelId,
			SessionId: session.SessionId,
			IsActive: session.IsActive,
			JoinAt: session.JoinAt,
			CreatedAt: session.CreatedAt,
		})

	}
}