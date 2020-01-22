package db

import (
	"github.com/pashkapo/catalog-lite/core"
	"github.com/pashkapo/catalog-lite/models"
)

func (db *Database) GetBuildings(page, count int) ([]*models.Building, error) {
	if page == 0 {
		page = core.DefaultPage
	}
	if count == 0 {
		count = core.DefaultCount
	}

	offset := count * (page - 1)

	rows, err := db.Query("SELECT id, country, city, street, house FROM buildings ORDER BY id OFFSET $1 LIMIT $2", offset, count)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	buildings := make([]*models.Building, 0)
	for rows.Next() {
		building := new(models.Building)
		err := rows.Scan(&building.Id, &building.Country, &building.City, &building.Street, &building.House)
		if err != nil {
			return nil, err
		}
		buildings = append(buildings, building)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return buildings, nil
}
