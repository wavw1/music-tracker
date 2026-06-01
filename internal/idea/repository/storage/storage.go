package storage

import (
	"sync"

	"github.com/wavw1/music-tracker/internal/model"
)

type StorageRepo struct {
	Ideas map[int]model.Idea
	Index int
	Mu    *sync.RWMutex
}

func NewStorageRepo(
	ideas map[int]model.Idea, mu *sync.RWMutex,
) StorageRepo {
	return StorageRepo{
		Ideas: ideas,
		Index: 1,
		Mu:    mu,
	}
}
