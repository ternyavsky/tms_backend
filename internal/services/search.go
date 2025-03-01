package services

import (
	"github.com/ternyavsky/tms_backend/internal/dto"
	"github.com/ternyavsky/tms_backend/internal/repo"
)

type SearchService struct {
	linkRepository *repo.LinkRepository
}

func NewSearchService(linkRepository *repo.LinkRepository) *SearchService {
	return &SearchService{
		linkRepository: linkRepository,
	}
}

func (context SearchService) SearchChannel(channelUrl string) (dto.LinkResponseDto, error) {
	value, err := context.linkRepository.CheckLinkByChannelUrl(channelUrl)
	return dto.CreateLinkResponseDto(value), err
}

func (context SearchService) GetSearchStatus(linkId uint) (dto.LinkResponseDto, error) {
	value, err := context.linkRepository.CheckLinkStatus(linkId)
	return dto.CreateLinkResponseDto(value), err
}
