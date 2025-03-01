package repo

import (
	"errors"
	"time"

	"github.com/ternyavsky/tms_backend/internal/domain"
	"github.com/ternyavsky/tms_backend/internal/dto"
	"gorm.io/gorm"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (context *SessionRepository) GetAllSessions() ([]domain.Session, error) {
	var sessions []domain.Session

	err := context.db.Model(&domain.Session{}).Preload("Source").Find(&sessions).Error
	if err != nil {
		return []domain.Session{}, errors.New("failed to get sessions")
	}

	return sessions, nil
}

func (context *SessionRepository) CreateSession(createSessionRequestDto dto.SessionRequestDto) (domain.Session, error) {
	session := domain.Session{
		ChannelId: createSessionRequestDto.ChannelId,
		SessionId: createSessionRequestDto.SessionId,
		IsActive:  createSessionRequestDto.IsActive,
		SourceId:  createSessionRequestDto.SourceId,
		JoinAt:    time.Now(),
	}

	var source domain.Source
	if err := context.db.First(&source, createSessionRequestDto.SourceId).Error; err != nil {
		return domain.Session{}, errors.New("source not found")
	}
	result := context.db.Create(&session)
	if result.Error != nil {
		return domain.Session{}, errors.New("failed to create session")
	}

	return session, nil
}
