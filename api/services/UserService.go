package services

import "github.com/inversion-go/alejoab12/user-api/api/models"

func FindById(id string) models.User {
	return models.User{
		Document:       "1010194766",
		DocumentTypeId: models.DocumentType{Id: "1", Name: "Cedula"},
		Name:           "Manuel Alejandro",
		LastName:       "Alcala Bustos",
		Id:             "34234234",
		Email:          "alejoab12@hotmail.com"}
}
