package controllers

import "github.com/ternyavsky/tms_backend/internal/services"

type SearchController struct {
	linkService services.LinkService
}

func NewLinkController(linkService services.LinkService) *LinkController {
	return &LinkController{linkService: linkService}
}
