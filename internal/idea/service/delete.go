package idea_service

import (
	"context"

	core_errors "github.com/wavw1/music-tracker/internal/errors"
)

func (s *IdeaServiceStruct) DeleteIdea(
	ctx context.Context,
	ideaID int,
	userID int64,
) error {
	idea, err := s.repo.GetIdea(
		ctx,
		ideaID,
	)
	if err != nil {
		return err
	}

	if idea.UserID != userID {
		return core_errors.ErrForbidden
	}

	err = s.repo.DeleteIdea(
		ctx,
		ideaID,
	)
	if err != nil {
		return err
	}

	return nil
}
