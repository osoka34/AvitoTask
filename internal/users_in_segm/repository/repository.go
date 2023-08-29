package repository

import (
	"AvitoTask/internal/s_constant"
	"AvitoTask/internal/users_in_segm"
	"AvitoTask/pkg/utils"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jmoiron/sqlx"
)

type postgresRepository struct {
	poolDb *pgxpool.Pool
	db     *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB, poolDb *pgxpool.Pool) users_in_segm.Repository {
	return &postgresRepository{db: db, poolDb: poolDb}
}

func (p *postgresRepository) InsertUserInSegments(params *users_in_segm.InsertUserInSegParams) error {
	var (
		query = `
			INSERT INTO %[1]s 
				(user_id, segment_id, ttl)
			VALUES %[2]s;
		`
		values []interface{} = append([]interface{}{params.UserId}, append(params.SegmentNames, params.Ttl...)...)
	)
	query = fmt.Sprintf(query, s_constant.UsersInSegment,
		utils.StringSliceToDollarPsqlArrayWithTTL(len(params.SegmentNames), params.Ttl))

	con, err := p.db.Query(query, values...)
	if err != nil {
		return err
	}

	return utils.CloseConnection(con)
}

func (p *postgresRepository) DeleteUserFromSeg(params *users_in_segm.DeleteUserFromSegParams) error {
	var (
		query = `
			DELETE FROM %[1]s
				WHERE user_id = $1 AND
      				segment_id IN (SELECT segment_id FROM public.segment
						WHERE segment_name IN %[2]s);
		`
		values []interface{} = append([]interface{}{params.UserId}, params.SegmentNames...)
	)

	query = fmt.Sprintf(query, s_constant.UsersInSegment, utils.StringSliceToDollarPsqlArray(len(params.SegmentNames), 1))

	con, err := p.db.Query(query, values...)
	if err != nil {
		return err
	}

	return utils.CloseConnection(con)
}

func (p *postgresRepository) GetAllSegByUserId(params *users_in_segm.SelectBy) ([]string, error) {

	var (
		dataRaw []string
		query   = `
			SELECT segment_name FROM %[1]s WHERE segment_id IN
			(SELECT segment_id FROM %[2]s WHERE user_id = $1);
		`
		values []interface{} = []interface{}{
			params.UserId,
		}
	)

	query = fmt.Sprintf(query, s_constant.SegmentDB, s_constant.UsersInSegment)
	if err := p.db.Select(&dataRaw, query, values...); err != nil {
		return dataRaw, err
	}

	if len(dataRaw) == 0 {
		return dataRaw, fmt.Errorf("no data in %[1]s table", s_constant.UsersInSegment)
	}

	return dataRaw, nil
}

func (p *postgresRepository) DeleteByTtl() error {
	var (
		query = `
		WITH deleted_rows AS (
			DELETE FROM %[1]s
			WHERE ttl <= $1
			RETURNING user_id, segment_id, ttl
		)
		INSERT INTO %[2]s (user_id, segment_name, date_event, in_event)
		SELECT d.user_id, s.segment_name, $1, false
		FROM deleted_rows d
		JOIN public.segment s ON d.segment_id = s.segment_id;
			`

		values []any = []any{
			utils.GetMoscowTime(),
		}
	)

	query = fmt.Sprintf(query, s_constant.UsersInSegment, s_constant.StatisticsDB)

	con, err := p.db.Query(query, values...)
	if err != nil {
		return err
	}

	return utils.CloseConnection(con)
}

func (p *postgresRepository) CountUsersWithExpiredTtl() (int, error) {
	var (
		count int
		query = `
			SELECT COALESCE(COUNT(ttl),0) as count FROM %[1]s WHERE ttl <= $1;
		`
		values []any = []any{
			utils.GetMoscowTime(),
		}
	)

	query = fmt.Sprintf(query, s_constant.UsersInSegment)
	if err := p.db.QueryRow(query, values...).Scan(&count); err != nil {
		return count, err
	}

	return count, nil
}
