package services

import (
	"github.com/ternyavsky/tms_backend/internal/repo"
	"github.com/ternyavsky/tms_backend/internal/responses"
)

type SessionService struct{
	sessionRepository *repo.SessionRepository
}

func (context SessionService) GetAllSessions() ([]responses.SessionResponse, error){
	values, err := context.sessionRepository.GetAllSessions()
	return responses.GetAllSessionsResponse(values), err
}
