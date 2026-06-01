package idea_service

import (
	"context"

	idea_domain "github.com/wavw1/music-tracker/internal/idea/domain"
	idea_repository "github.com/wavw1/music-tracker/internal/idea/repository"
	"github.com/wavw1/music-tracker/internal/model"
)

type IdeaServiceStruct struct {
	repo idea_repository.IdeaRepository
}

func NewIdeaServiceStruct(
	repo idea_repository.IdeaRepository,
) *IdeaServiceStruct {
	return &IdeaServiceStruct{
		repo: repo,
	}
}

type IdeaService interface {
	CreateIdea(
		ctx context.Context,
		ideaRequest model.Idea,
		userID int64,
	) (model.Idea, error)

	GetIdeas(
		ctx context.Context,
		limit int,
		offset int,
		userID int64,
	) ([]model.Idea, error)

	GetIdea(
		ctx context.Context,
		idParam int,
		userID int64,
	) (model.Idea, error)

	UpdateIdea(
		ctx context.Context,
		ideaID int,
		idea idea_domain.IdeaServiceInput,
		userID int64,
	) (
		model.Idea,
		error,
	)

	DeleteIdea(
		ctx context.Context,
		ideaID int,
		userID int64,
	) error
}
