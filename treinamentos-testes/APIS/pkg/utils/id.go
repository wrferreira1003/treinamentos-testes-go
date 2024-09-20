package utils

import "github.com/google/uuid"

type ID = uuid.UUID

// Criando um novo ID
func NewID() ID {
	return ID(uuid.New())
}

// Verifico se o ID é válido
func ParseID(id string) (ID, error) {
	idParsed, err := uuid.Parse(id)
	return ID(idParsed), err
}
