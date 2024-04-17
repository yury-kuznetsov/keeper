package data

import (
	"gophkeeper/internal/client/model"
)

// FindAll retrieves all the data records from the repository.
//
// It executes a SELECT query to fetch all the data records from the "data" table.
// The retrieved data is then mapped to an array of model.Data structs.
// The resulting array and any errors encountered during the execution are returned.
func (r *Repository) FindAll() ([]model.Data, error) {
	rows, err := r.db.Query("SELECT * FROM data")
	if err != nil {
		return nil, err
	}

	var entities []model.Data

	for rows.Next() {
		var entity model.Data
		err := rows.Scan(
			&entity.ID,
			&entity.Category,
			&entity.Data,
			&entity.Version,
		)
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return entities, nil
}
