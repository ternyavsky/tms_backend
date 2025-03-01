package services

import (
	"github.com/ternyavsky/tms_backend/internal/dto"
	"github.com/ternyavsky/tms_backend/internal/repo"
)

type LinkService struct {
	linkRepository *repo.LinkRepository
}

func NewLinkService(linkRepository *repo.LinkRepository) *LinkService {
	return &LinkService{
		linkRepository: linkRepository,
	}
}

func (context LinkService) GetAllLinks() ([]dto.LinkResponseDto, error) {
	values, err := context.linkRepository.GetAllLinks()
	return dto.GetAllLinksResponseDto(values), err
}

func (context LinkService) CreateLink(createLinkRequestDto dto.LinkRequestDto) (dto.LinkResponseDto, error) {
	value, err := context.linkRepository.CreateLink(createLinkRequestDto)
	return dto.CreateLinkResponseDto(value), err
}
