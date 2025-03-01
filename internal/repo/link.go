package repo

import (
	"errors"

	"github.com/ternyavsky/tms_backend/internal/domain"
	"github.com/ternyavsky/tms_backend/internal/dto"
	"gorm.io/gorm"
)

type LinkRepository struct {
	db *gorm.DB
}

func NewLinkRepository(db *gorm.DB) *LinkRepository {
	return &LinkRepository{db: db}
}

func (context *LinkRepository) GetAllLinks() ([]domain.Link, error) {
	var links []domain.Link

	err := context.db.Model(&domain.Link{}).Find(&links).Error
	if err != nil {
		return []domain.Link{}, errors.New("failed to get links")
	}
	return links, nil
}

func (context *LinkRepository) CheckLinkByChannelUrl(channelUrl string) (domain.Link, error) {
	var link domain.Link
	err := context.db.Where("link_value = ?", channelUrl).First(&link).Error
	if err != nil {
		return link, errors.New("failed to get link by this channelUrl")
	}
	link = domain.Link{
		LinkValue: channelUrl,
		LinkType:  domain.InvalidRequest,
		IsValid:   false,
	}
	err = context.db.Create(&link).Error
	if err != nil {
		return link, errors.New("failed to create link")
	}
	return link, nil
}

func (context *LinkRepository) CheckLinkStatus(linkId uint) (domain.Link, error) {
	var link domain.Link
	err := context.db.First(&link, "id = ?", linkId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Link{}, errors.New("link not found")
		}
		return domain.Link{}, errors.New("failed to get link")
	}
	return link, nil
}

func (context *LinkRepository) CreateLink(createLinkRequestDto dto.LinkRequestDto) (domain.Link, error) {
	link := domain.Link{
		LinkValue: createLinkRequestDto.LinkValue,
		LinkType:  createLinkRequestDto.LinkType,
		IsValid:   createLinkRequestDto.IsValid,
	}

	// var source domain.Source
	// if err := context.db.First(&source, createSessionRequestDto.SourceId).Error; err != nil {
	// 	return domain.Session{}, errors.New("source not found")
	// }
	result := context.db.Create(&link)
	if result.Error != nil {
		return domain.Link{}, errors.New("failed to create link")
	}

	return link, nil
}
