package idea_handler

import idea_service "github.com/wavw1/music-tracker/internal/idea/service"

type IdeaHandler struct {
	service idea_service.IdeaService
}

func NewIdeaHandler(s idea_service.IdeaService) *IdeaHandler {
	return &IdeaHandler{
		service: s,
	}
}
