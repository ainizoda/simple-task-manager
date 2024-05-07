package task

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

const minTitleLength = 4
const minDescriptionLength = 16

var (
	ErrTitleShort       = fmt.Errorf("error: title must contain at least %d characters", minTitleLength)
	ErrDescriptionShort = fmt.Errorf("error: description must contain at least %d characters", minDescriptionLength)
	ErrTaskNotFound     = fmt.Errorf("error: task not found")
)

var Statuses = []TaskStatus{Done, InProgress, New}

type Service struct {
	tasks []*Task
}

func (s *Service) findTaskById(id string) (*Task, error) {
	for _, t := range s.tasks {
		if t.ID == id {
			return t, nil
		}
	}
	return nil, ErrTaskNotFound
}

func (s *Service) Create(title string, description string) error {
	if len(strings.TrimSpace(title)) < minTitleLength {
		return ErrTitleShort
	}
	if len(strings.TrimSpace(description)) < minDescriptionLength {
		return ErrDescriptionShort
	}
	task := &Task{
		ID:          uuid.NewString(),
		Status:      New,
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
	}
	s.tasks = append(s.tasks, task)
	return nil
}

func (s *Service) GetAll() []*Task {
	return s.tasks
}

func (s *Service) Update(updatedTask *Task) error {
	task, err := s.findTaskById(updatedTask.ID)
	if err != nil {
		return err
	}

	*task = *updatedTask
	return nil
}

func (s *Service) RemoveById(id string) error {
	task, err := s.findTaskById(id)
	if err != nil {
		return err
	}
	var updatedList []*Task
	for _, t := range s.tasks {
		if t.ID != task.ID {
			updatedList = append(updatedList, t)
		}
	}
	s.tasks = updatedList
	return nil
}
