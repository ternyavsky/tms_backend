package repo

import (
	"errors"

	"github.com/ternyavsky/tms_backend/internal/domain"
	"gorm.io/gorm"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepository{
	return &SessionRepository{db: db}
}

func (context SessionRepository) GetAllSessions() ([]domain.Session, error) {

	var sessions []domain.Session

	err := context.db.Model(&domain.Session{}).Find(&sessions).Error

	if err != nil {
		return []domain.Session{}, errors.New("failed to get sessions")
	}

	return sessions, nil
}