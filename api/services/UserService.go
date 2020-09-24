package services

import (
	"github.com/google/uuid"
	"github.com/inventario-go/alejoab12/user-api/api/models"
	"github.com/inventario-go/alejoab12/user-api/api/repository"
	"time"
)

func FindById(id string) models.User {
	return repository.FindUserById(id)
}
func CreateUser(user *models.User) {
	user.CreatedAt = time.Now()
	uuid, _ := uuid.NewRandom()
	user.ID = uuid.String()
	repository.CreateUser(user)
}
