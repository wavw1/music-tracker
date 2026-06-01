package idea_service

import (
	"context"
	"fmt"

	core_errors "github.com/wavw1/music-tracker/internal/errors"
	"github.com/wavw1/music-tracker/internal/model"
)

func (s *IdeaServiceStruct) CreateIdea(
	ctx context.Context,
	ideaRequest model.Idea,
	userID int64,
) (model.Idea, error) {
	var key string = ideaRequest.Key

	if ideaRequest.Bpm <= 50 || ideaRequest.Bpm >= 550 {
		return model.Idea{}, core_errors.ErrInvalidBPM
	}

	if ideaRequest.Key == "" {
		key = "Unknown key"
	}

	ideaRequest.UserID = userID

	serviceModel, err := s.repo.CreateIdea(ctx, model.Idea{
		ID:     ideaRequest.ID,
		UserID: ideaRequest.UserID,
		Title:  ideaRequest.Title,
		Bpm:    ideaRequest.Bpm,
		Key:    key,
		Status: "draft",
		Tags:   []string{"beat", "draft"},
	})
	if err != nil {
		return model.Idea{}, fmt.Errorf("create idea repository error: %w", err)
	}

	return serviceModel, nil
}
