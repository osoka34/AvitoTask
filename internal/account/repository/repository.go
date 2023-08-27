package repository

import (
	"AvitoTask/internal/account"
	"AvitoTask/internal/s_constant"
	"AvitoTask/pkg/utils"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jmoiron/sqlx"
)

type postgresRepository struct {
	poolDb *pgxpool.Pool
	db     *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB, poolDb *pgxpool.Pool) account.Repository {
	return &postgresRepository{db: db, poolDb: poolDb}
}

func (p *postgresRepository) CreateAccount(params *account.CreateUserParams) (uint64, error) {
	var (
		userId uint64

		query string = `
			INSERT INTO %[1]s
				(user_id, user_name)
			VALUES
				($1, $2)
			RETURNING
				user_id;
		`
		values []any = []any{
			params.UserId, params.UserName,
		}
	)

	query = fmt.Sprintf(query, s_constant.UserDB)

	if err := p.db.QueryRow(query, values...).Scan(&userId); err != nil {
		return userId, err
	}

	return userId, nil
}

func (p *postgresRepository) DeleteAccount(params *account.DeleteUserParams) error {
	var (
		query = `
			DELETE FROM %[1]s
			WHERE user_id = $1;
			`
		values []any = []any{
			params.UserId,
		}
	)
	query = fmt.Sprintf(query, s_constant.UserDB)

	con, err := p.db.Query(query, values...)
	if err != nil {
		return err
	}

	return utils.CloseConnection(con)
}
