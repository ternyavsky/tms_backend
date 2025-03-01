package services

import (
	"github.com/ternyavsky/tms_backend/internal/dto"
	"github.com/ternyavsky/tms_backend/internal/repo"
)

type SessionService struct {
	sessionRepository *repo.SessionRepository
}

func NewSessionService(sessionRepository *repo.SessionRepository) *SessionService {
	return &SessionService{
		sessionRepository: sessionRepository,
	}
}

func (context SessionService) GetAllSessions() ([]dto.SessionResponseDto, error) {
	values, err := context.sessionRepository.GetAllSessions()
	return dto.GetAllSessionsResponseDto(values), err
}

func (context SessionService) CreateSession(createSessionRequestDto dto.SessionRequestDto) (dto.SessionResponseDto, error) {
	value, err := context.sessionRepository.CreateSession(createSessionRequestDto)
	return dto.CreateSessionResponseDto(value), err
}
