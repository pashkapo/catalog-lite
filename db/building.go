package db

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/pashkapo/catalog-lite/config"
	"github.com/pashkapo/catalog-lite/model"
)

func (db *Database) GetBuildings(page, count int) ([]*model.Building, error) {
	if page == 0 {
		page = config.DefaultPage
	}
	if count == 0 {
		count = config.DefaultCount
	}

	offset := count * (page - 1)

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	firmsQuery := psql.Select("b.id, b.country, b.city, b.street, b.house, b.location[0], b.location[1]").
		From("buildings b").
		OrderBy("id").
		Offset(uint64(offset)).
		Limit(uint64(count))

	sql, args, err := firmsQuery.ToSql()

	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	buildings := make([]*model.Building, 0)
	for rows.Next() {
		building := new(model.Building)
		err := rows.Scan(&building.Id, &building.Country, &building.City, &building.Street, &building.House, &building.Location.Long, &building.Location.Lat)
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
