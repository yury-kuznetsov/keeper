package data

import (
	"gophkeeper/internal/client/model"

	"github.com/google/uuid"
)

// Status retrieves and compares data from the local and remote databases, and returns a list
// of model.DataVersion objects representing the status of each data item.
//
// Returns:
// - []model.DataVersion: a list of data versions representing the status of each data item
// - error: an error object if there was any issue retrieving or comparing the data
func (s *Service) Status() ([]model.DataVersion, error) {
	// извлекаем записи из локальной базы
	dataLocal, err := s.r.FindAll()
	if err != nil {
		return nil, err
	}

	// извлекаем записи из удаленной базы
	dataRemote, err := s.c.Status()
	if err != nil {
		return nil, err
	}

	// создаем словарь для быстрого доступа к записям
	dataMap := make(map[uuid.UUID]model.DataVersion)
	for _, data := range dataLocal {
		dataMap[data.ID] = model.DataVersion{
			ID:            data.ID,
			VersionLocal:  data.Version,
			VersionRemote: 0,
		}
	}

	// проверяем каждую запись из удаленной базы
	for _, remote := range dataRemote {
		local, ok := dataMap[remote.ID]
		if ok {
			local.VersionRemote = remote.Version
			dataMap[remote.ID] = local
		} else {
			dataMap[remote.ID] = model.DataVersion{
				ID:            remote.ID,
				VersionLocal:  0,
				VersionRemote: remote.Version,
			}
		}
	}

	// создаем список для хранения версий
	var versions []model.DataVersion
	for _, data := range dataMap {
		versions = append(versions, data)
	}

	return versions, nil
}
