package repository

import (
	"AvitoTask/internal/s_constant"
	"AvitoTask/internal/statistics"
	"AvitoTask/pkg/utils"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jmoiron/sqlx"
)

type postgresRepository struct {
	poolDb *pgxpool.Pool
	db     *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB, poolDb *pgxpool.Pool) statistics.Repository {
	return &postgresRepository{db: db, poolDb: poolDb}
}

func (p *postgresRepository) AddRows(params *statistics.InsertParams) error {
	var (
		query = `
		INSERT INTO %[1]s
		(user_id, segment_name, date_event, in_event)
		VALUES %[2]s;
		`

		values []interface{} = append([]interface{}{
			params.UserId, params.Time, params.In,
		}, params.Segment_names...)
	)

	query = fmt.Sprintf(query, s_constant.StatisticsDB, utils.InserrtRowsDollarParams(len(params.Segment_names)))
	con, err := p.db.Query(query, values...)
	if err != nil {
		return err
	}

	return utils.CloseConnection(con)
}

func (p *postgresRepository) SelectByDates(params *statistics.SelectParams) ([]statistics.SelectOut, error) {
	var (
		data  []statistics.SelectOut
		query = `
			SELECT user_id, segment_name, date_event, in_event 
			FROM %[1]s 
			WHERE date_event BETWEEN $1 and $2 ORDER BY user_id;
		`
		values []interface{} = []interface{}{
			params.DateFrom, params.DateTo,
		}
	)
	query = fmt.Sprintf(query, s_constant.StatisticsDB)

	if err := p.db.Select(&data, query, values...); err != nil {
		return data, err
	}

	if len(data) == 0 {
		return data, fmt.Errorf("no data in %[1]s table", s_constant.StatisticsDB)
	}

	return data, nil
}
