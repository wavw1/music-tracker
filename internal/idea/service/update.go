package idea_service

import (
	"context"
	"fmt"

	core_errors "github.com/wavw1/music-tracker/internal/errors"
	idea_domain "github.com/wavw1/music-tracker/internal/idea/domain"
	"github.com/wavw1/music-tracker/internal/model"
)

func (s *IdeaServiceStruct) UpdateIdea(
	ctx context.Context,
	ideaID int,
	idea idea_domain.IdeaServiceInput,
	userID int64,
) (
	model.Idea,
	error,
) {
	var patch idea_domain.IdeaPatch

	repoIdea, err := s.repo.GetIdea(
		ctx,
		ideaID,
	)
	if err != nil {
		return model.Idea{}, err
	}

	if idea.Title != nil {
		patch.Title = idea.Title
	}

	if idea.Key != nil {
		if *idea.Key == "" {
			unknownKey := "Unknown key"

			patch.Key = &unknownKey
		} else {
			patch.Key = idea.Key
		}
	}

	if idea.Bpm != nil {
		if *idea.Bpm <= 50 || *idea.Bpm >= 550 {
			return model.Idea{}, core_errors.ErrInvalidBPM
		}

		patch.Bpm = idea.Bpm
	}

	if repoIdea.UserID != userID {
		return model.Idea{}, core_errors.ErrForbidden
	}

	ideaResponse, err := s.repo.UpdateIdea(
		ctx,
		ideaID,
		patch,
	)
	if err != nil {
		return model.Idea{}, fmt.Errorf("failed to update idea: %w", err)
	}

	return ideaResponse, nil
}
