package idea_service

import (
	"context"
	"fmt"

	"github.com/wavw1/music-tracker/internal/model"
)

func (s *IdeaServiceStruct) GetIdeas(
	ctx context.Context,
	limit int,
	offset int,
	userID int64,
) ([]model.Idea, error) {
	ideas, err := s.repo.GetIdeas(
		ctx,
		limit,
		offset,
		userID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to save ideas. repository error: %w", err)
	}

	return ideas, nil
}
