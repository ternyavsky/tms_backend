package dto

import (
	"time"

	"github.com/ternyavsky/tms_backend/internal/domain"
)

type LinkRequestDto struct {
	LinkValue string          `json:"link_value"`
	LinkType  domain.LinkType `json:"link_type"`
	IsValid   bool            `json:"is_valid"`
}
type LinkResponseDto struct {
	Id        uint            `json:"id"`
	LinkValue string          `json:"link_value"`
	LinkType  domain.LinkType `json:"link_type"`
	IsValid   bool            `json:"is_valid"`
	CreatedAt time.Time       `json:"created_at"`
	Sources   []domain.Source
}

func GetAllLinksResponseDto(links []domain.Link) []LinkResponseDto {
	var values []LinkResponseDto
	for _, link := range links {
		values = append(values, LinkResponseDto{
			Id:        link.Id,
			LinkValue: link.LinkValue,
			LinkType:  link.LinkType,
			IsValid:   link.IsValid,
			Sources:   link.Sources,
			CreatedAt: link.CreatedAt,
		})
	}
	return values
}

func CreateLinkResponseDto(link domain.Link) LinkResponseDto {
	return LinkResponseDto{
		Id:        link.Id,
		LinkValue: link.LinkValue,
		LinkType:  link.LinkType,
		IsValid:   link.IsValid,
		CreatedAt: link.CreatedAt,
		Sources:   link.Sources,
	}
}
