package storage

import (
	"context"
	"github.com/datatug/datatug/packages/models"
)

// Store defines interface for loading & saving DataTug projects
type Store interface {
	// GetProjects returns list of projects
	GetProjects(ctx context.Context) (projectBriefs []models.ProjectBrief, err error)
	// Project returns project store
	Project(id string) ProjectStore
}

var _ Store = (*NoOpStore)(nil)

type NoOpStore struct {
}

func (n NoOpStore) GetProjects(ctx context.Context) (projectBriefs []models.ProjectBrief, err error) {
	panic("implement me")
}

func (n NoOpStore) Project(id string) ProjectStore {
	panic("implement me")
}