package repository

import (
	"AvitoTask/internal/s_constant"
	"AvitoTask/internal/segment"
	"AvitoTask/pkg/utils"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jmoiron/sqlx"
)

type postgresRepository struct {
	poolDb *pgxpool.Pool
	db     *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB, poolDb *pgxpool.Pool) segment.Repository {
	return &postgresRepository{db: db, poolDb: poolDb}
}

func (p *postgresRepository) CreateSegment(params *segment.CreateSegmentParams) (string, error) {
	var (
		segmentName string

		query string = `
			INSERT INTO %[1]s
				(segment_name)
			VALUES
				($1)
			RETURNING
				segment_name;
		`
		values []any = []any{
			params.SegmentName,
		}
	)

	query = fmt.Sprintf(query, s_constant.SegmentDB)

	if err := p.db.QueryRow(query, values...).Scan(&segmentName); err != nil {
		return segmentName, err
	}

	return segmentName, nil
}

func (p *postgresRepository) UpdateSegmentName(params *segment.UpdateSegmentParams) error {
	var (
		query string = `
			UPDATE
				%[1]s
			SET
				segment_name = $2
			WHERE
				segment_id = $1;
		`
		values []any = []any{
			params.SegmentId, params.NewSegName,
		}
	)

	query = fmt.Sprintf(query, s_constant.SegmentDB)

	con, err := p.db.Query(query, values...)
	if err != nil {
		return err
	}

	return utils.CloseConnection(con)
}

func (p *postgresRepository) DeleteSegmentByName(params *segment.DeleteSegmParams) error {
	var (
		query = `
			DELETE FROM %[1]s
			WHERE segment_name = $1;
			`
		values []any = []any{
			params.SegmentName,
		}
	)
	query = fmt.Sprintf(query, s_constant.SegmentDB)

	con, err := p.db.Query(query, values...)
	if err != nil {
		return err
	}

	return utils.CloseConnection(con)
}

func (p *postgresRepository) GetAmountInSegment(params *segment.GetSegmentParams) (int, error) {
	var (
		count int
		query = `
			SELECT COALESCE(count(*),0) AS count FROM %[1]s WHERE %[2]s=$1;
			`
		values []any = []any{
			params.Value,
		}
	)
	query = fmt.Sprintf(query, s_constant.SegmentDB, params.FieldName)

	if err := p.db.QueryRow(query, values...).Scan(&count); err != nil {
		return count, err
	}

	return count, nil
}

func (p *postgresRepository) CreateSegmentWithAutoAdd(params *segment.CreateSegmentParams) error {
	var (
		query = `
			WITH created AS ( INSERT INTO %[1]s -- segment
			(segment_name) VALUES ($1) RETURNING segment_id), -- segment_name
			inserted AS (
				INSERT INTO %[2]s (user_id, segment_id) -- users_in_segm
				SELECT u.user_id, (SELECT segment_id FROM created)
				FROM public.user u
					WHERE random() < $2 -- probability
			   RETURNING user_id, segment_id
			)
			INSERT INTO %[3]s (user_id, segment_name, date_event, in_event) -- statistics
			SELECT i.user_id, $1, CURRENT_DATE, true  -- segment_name
			FROM inserted i;
			`
		values []any = []any{
			params.SegmentName, params.Probability,
		}
	)

	query = fmt.Sprintf(query, s_constant.SegmentDB, s_constant.UsersInSegment, s_constant.StatisticsDB)

	con, err := p.db.Query(query, values...)
	if err != nil {
		return err
	}

	return utils.CloseConnection(con)
}
