package data

import (
	"errors"

	"github.com/google/uuid"
)

// Push sends a model to the server and updates its version if there were changes.
// It returns an error if the entity is not found, if there was an issue sending the data to the server,
// or if there was an issue saving the updated version.
func (s *Service) Push(id uuid.UUID) error {
	// ищем модель в базе
	entity := s.r.Find(id)
	if entity.ID == uuid.Nil {
		return errors.New("entity not found")
	}

	// отправляем на сервер
	version, err := s.c.Push(entity)
	if err != nil {
		return err
	}

	// проверяем наличие изменений
	if entity.Version == version {
		return nil
	}

	// меняем версию на серверную
	entity.Version = version

	return s.r.Save(entity)
}
