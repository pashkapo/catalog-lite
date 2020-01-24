package db

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/pashkapo/catalog-lite/config"
	"github.com/pashkapo/catalog-lite/model"
)

func (db *Database) GetFirms(page, count int, filter *model.FirmFilter) ([]*model.Firm, error) {
	if page == 0 {
		page = config.DefaultPage
	}
	if count == 0 {
		count = config.DefaultCount
	}

	offset := count * (page - 1)

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	// При переносе на "прод" обнаружил проблему с неверным построением плана запроса в бд на heroku,
	// версия(11.6) там отличается от локальной(12.0). Что-то связано с order by и limit, пока нет времени разобраться
	// временно убрал order by
	firmsQuery := psql.Select("f.id, f.name, b.id, b.country, b.city, b.street, b.house, b.location[0], b.location[1]").
		From("firms f").
		Join("buildings b ON f.building_id = b.id").
		Offset(uint64(offset)).
		Limit(uint64(count))

	if filter.BuildingId != 0 {
		firmsQuery = firmsQuery.Where(sq.Eq{"f.building_id": filter.BuildingId})
	}

	if filter.RubricId != 0 {
		firmsQuery = firmsQuery.Join(`(
			with recursive r as (
				select id, parent_id
				from rubrics
				where id = ?
				union
				select rs.id, rs.parent_id
				from rubrics rs
					join r on rs.parent_id = r.id
			)
			select *
			from r
				join firms_rubrics fr on r.id = fr.rubric_id) rr on rr.firm_id = f.id`, filter.RubricId)
	}

	if filter.InRadius.Radius != 0 {
		firmsQuery = firmsQuery.Where(
			"st_dwithin(location::geometry::geography, st_makepoint(?, ?)::geography, ?)",
			filter.InRadius.Point.Long,
			filter.InRadius.Point.Lat,
			filter.InRadius.Radius,
		)
	}

	if filter.Search != "" {
		firmsQuery = firmsQuery.Where("to_tsvector('english',name) @@ plainto_tsquery(?)", filter.Search)
	}

	sql, args, err := firmsQuery.ToSql()

	//log.Info(filter.Search)
	//log.Info(args)
	//log.Info(sql)

	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	firms := make([]*model.Firm, 0)
	for rows.Next() {
		firm := new(model.Firm)
		err := rows.Scan(
			&firm.Id,
			&firm.Name,
			&firm.Building.Id,
			&firm.Building.Country,
			&firm.Building.City,
			&firm.Building.Street,
			&firm.Building.House,
			&firm.Building.Location.Long,
			&firm.Building.Location.Lat)
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

func (db *Database) GetFirmById(id uint) (*model.Firm, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	firmsQuery := psql.Select("f.id, f.name, b.id, b.country, b.city, b.street, b.house, b.location[0], b.location[1]").
		From("firms f").
		Join("buildings b ON f.building_id = b.id").
		Where(sq.Eq{"f.id": id})
	sql, args, err := firmsQuery.ToSql()

	var firm model.Firm

	err = db.QueryRow(sql, args...).
		Scan(
			&firm.Id,
			&firm.Name,
			&firm.Building.Id,
			&firm.Building.Country,
			&firm.Building.City,
			&firm.Building.Street,
			&firm.Building.House,
			&firm.Building.Location.Long,
			&firm.Building.Location.Lat)
	if err != nil {
		return nil, err
	}

	firm.PhoneNumbers, err = db.GetFirmPhoneNumbers(id)
	if err != nil {
		return nil, err
	}

	firm.Rubrics, err = db.GetFirmRubrics(id)
	if err != nil {
		return nil, err
	}

	return &firm, nil
}

func (db *Database) GetFirmPhoneNumbers(id uint) ([]string, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	firmsQuery := psql.Select("fpn.phone_number").
		From("firms_phone_numbers fpn").
		Where(sq.Eq{"fpn.firm_id": id})

	sql, args, err := firmsQuery.ToSql()
	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	phones := make([]string, 0)
	for rows.Next() {
		phone := ""
		err := rows.Scan(&phone)
		if err != nil {
			return nil, err
		}
		phones = append(phones, phone)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return phones, nil
}

func (db *Database) GetFirmRubrics(id uint) ([]*model.Rubric, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	firmsQuery := psql.Select("r.id, r.name").
		From("firms_rubrics fr").
		Join("rubrics r ON fr.rubric_id = r.id").
		Where(sq.Eq{"fr.firm_id": id})

	sql, args, err := firmsQuery.ToSql()

	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rubrics := make([]*model.Rubric, 0)
	for rows.Next() {
		rubric := new(model.Rubric)
		err := rows.Scan(&rubric.Id, &rubric.Name)
		if err != nil {
			return nil, err
		}
		rubrics = append(rubrics, rubric)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return rubrics, nil
}
