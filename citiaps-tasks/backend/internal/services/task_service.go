package services

import (
	"errors"
	"strings"
	"time"

	"citiaps-tasks-backend/internal/models"
	"citiaps-tasks-backend/internal/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	repo *repositories.TaskRepository
}

func NewTaskService(repo *repositories.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) Create(req models.CreateTaskRequest) (models.Task, error) {
	if strings.TrimSpace(req.Title) == "" {
		return models.Task{}, errors.New("el título es obligatorio")
	}

	now := time.Now()
	task := models.Task{
		Title:       strings.TrimSpace(req.Title),
		Description: strings.TrimSpace(req.Description),
		Completed:   false,
		Active:      true,
		IsDeleted:   false,
		Tags:        req.Tags,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	return s.repo.Create(task)
}

func (s *TaskService) FindAll() ([]models.Task, error) {
	return s.repo.FindAll()
}

func (s *TaskService) FindByID(id string) (*models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id inválido")
	}
	return s.repo.FindByID(objID)
}

func (s *TaskService) Complete(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("id inválido")
	}
	return s.repo.Complete(objID)
}

func (s *TaskService) Delete(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("id inválido")
	}
	return s.repo.SoftDelete(objID)
}

func (s *TaskService) Recover(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("id inválido")
	}
	return s.repo.Recover(objID)
}

func (s *TaskService) ToggleActive(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("id inválido")
	}
	return s.repo.ToggleActive(objID)
}

func (s *TaskService) DeletePermanent(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("id inválido")
	}
	return s.repo.DeletePermanent(objID)
}

func IsNotFound(err error) bool {
	return err == mongo.ErrNoDocuments
}
