package idea_repository

import (
	"context"

	idea_domain "github.com/wavw1/music-tracker/internal/idea/domain"
	"github.com/wavw1/music-tracker/internal/model"
)

type IdeaRepository interface {
	CreateIdea(
		ctx context.Context,
		ideaRequest model.Idea,
	) (model.Idea, error)

	GetIdeas(
		ctx context.Context,
		limit int,
		offset int,
		userID int64,
	) ([]model.Idea, error)

	GetIdea(
		ctx context.Context,
		id int,
	) (model.Idea, error)

	UpdateIdea(
		ctx context.Context,
		ideaID int,
		idea idea_domain.IdeaPatch,
	) (
		model.Idea,
		error,
	)

	DeleteIdea(
		ctx context.Context,
		ideaID int,
	) error
}
