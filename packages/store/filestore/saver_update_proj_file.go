package filestore

import (
	"github.com/datatug/datatug/packages/models"
)

func (s fileSystemSaver) loadProjectFile() (v models.ProjectFile, err error) {
	return LoadProjectFile(s.projDirPath)
}

func (s fileSystemSaver) updateProjectFileWithBoard(board models.Board) (err error) {
	projFile, err := s.loadProjectFile()
	if err != nil {
		return err
	}
	for _, b := range projFile.Boards {
		if b.ID == board.ID {
			if b.Title == board.Title {
				goto SAVED
			}
			b.Title = board.Title
			goto UPDATED
		}
	}
	projFile.Boards = append(projFile.Boards, &models.ProjBoardBrief{
		ProjectItem: models.ProjectItem{ID: board.ID, Title: board.Title},
		Parameters:  board.Parameters,
	})
UPDATED:
	err = s.putProjectFile(projFile)
SAVED:
	return err
}

func (s fileSystemSaver) updateProjectFile(updater func(projFile *models.ProjectFile) error) error {
	s.projFileMutex.Lock()
	defer func() {
		s.projFileMutex.Unlock()
	}()
	projFile, err := s.loadProjectFile()
	if err != nil {
		return err
	}
	if err = updater(&projFile); err != nil {
		return err
	}
	if err = s.putProjectFile(projFile); err != nil {
		return err
	}
	return nil
}

/*
func (s fileSystemSaver) updateProjectFileDeleteEntity(id string) (err error) {
	// Almost duplicates updateProjectFileDeleteBoard and other deletes
	projFile, err := s.loadProjectFile()
	if err != nil {
		return err
	}
	if len(projFile.Entities) == 0 {
		return
	}
	entities := make([]*models.ProjEntityBrief, len(projFile.Entities))
	for _, item := range projFile.Entities {
		if item.ID == id {
			continue
		}
	}
	if len(projFile.Entities) > len(entities) {
		projFile.Entities = entities
		err = s.putProjectFile(projFile)
	}
	return err
}

func (s fileSystemSaver) updateProjectFileDeleteBoard(id string) (err error) {
	// Almost duplicates updateProjectFileDeleteEntity and other deletes
	projFile, err := s.loadProjectFile()
	if err != nil {
		return err
	}
	if len(projFile.Entities) == 0 {
		return
	}
	items := make([]*models.ProjBoardBrief, len(projFile.Boards))
	for _, item := range projFile.Boards {
		if item.ID == id {
			continue
		}
	}
	if len(projFile.Boards) > len(items) {
		projFile.Boards = items
		err = s.putProjectFile(projFile)
	}
	return err
}
*/
