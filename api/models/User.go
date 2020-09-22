package models

type User struct {
	Id             string       `json:"id"`
	DocumentTypeId DocumentType `json:"documentTypeId"`
	Document       string       `json:document`
	Name           string       `json:"name"`
	LastName       string       `json:"lastName"`
	Email          string       `json:"json:email"`

}
