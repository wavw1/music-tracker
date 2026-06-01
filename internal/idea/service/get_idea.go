package idea_service

import (
	"context"

	core_errors "github.com/wavw1/music-tracker/internal/errors"
	"github.com/wavw1/music-tracker/internal/model"
)

func (s *IdeaServiceStruct) GetIdea(
	ctx context.Context,
	id int,
	userID int64,
) (
	model.Idea, error,
) {
	idea, err := s.repo.GetIdea(
		ctx,
		id,
	)
	if err != nil {
		return model.Idea{}, err
	}

	if idea.UserID != userID {
		return model.Idea{}, core_errors.ErrForbidden
	}

	return idea, nil
}
