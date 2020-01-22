package db

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/labstack/gommon/log"
	"github.com/pashkapo/catalog-lite/core"
	"github.com/pashkapo/catalog-lite/models"
)

func (db *Database) GetFirms(page, count int, filter models.FirmFilter) ([]*models.Firm, error) {
	if page == 0 {
		page = core.DefaultPage
	}
	if count == 0 {
		count = core.DefaultCount
	}

	offset := count * (page - 1)

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	firmsQuery := psql.Select("f.*").From("firms f").OrderBy("id").Offset(uint64(offset)).Limit(uint64(count))

	if filter.BuildingId != 0 {
		firmsQuery = firmsQuery.Where(sq.Eq{"building_id": filter.BuildingId})
	}

	if filter.RubricId != 0 {
		firmsQuery = firmsQuery.Join("firms_rubrics fr ON f.id = fr.firm_id").Where(sq.Eq{"rubric_id": filter.RubricId})
	}

	if filter.InRadius != 0 {
		firmsQuery = firmsQuery.Join("buildings b ON f.building_id = b.id").Where("st_dwithin(location::geometry::geography, st_makepoint(167.561104, 71.509529)::geography, 2000000)")
	}

	sql, args, err := firmsQuery.ToSql()

	log.Info(sql)

	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	firms := make([]*models.Firm, 0)
	for rows.Next() {
		firm := new(models.Firm)
		err := rows.Scan(&firm.Id, &firm.Name, &firm.BuildingId)
		if err != nil {
			return nil, err
		}
		firms = append(firms, firm)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return firms, nil
}

func (db *Database) GetFirmById(id uint) (*models.Firm, error) {
	var firm models.Firm

	err := db.QueryRow("SELECT id, name, building_id FROM firms where id = $1", id).Scan(&firm.Id, &firm.Name, &firm.BuildingId)
	if err != nil {
		return nil, err
	}

	return &firm, nil
}
