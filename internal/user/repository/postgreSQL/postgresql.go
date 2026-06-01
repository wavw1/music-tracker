package user_postgresql

import "github.com/jackc/pgx/v5/pgxpool"

type PostgreSql struct {
	pool *pgxpool.Pool
}

func NewPostgreSql(pool *pgxpool.Pool) *PostgreSql {
	return &PostgreSql{
		pool: pool,
	}
}
