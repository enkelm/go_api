package db

import (
	"context"

	"github.com/enkelm/go_api/db/model"

	"github.com/jackc/pgx/v5"
)

func (pg *postgres) GetUsers(ctx context.Context) ([]model.User, error) {
	query := `SELECT * FROM public.user`
	rows, err := pg.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows[model.User](rows, pgx.RowToStructByName)
}

func (pg *postgres) GetUserById(id string) (model.User, error) {
	var user model.User
	row := pg.db.QueryRow(context.Background(), "select * from public.user where id=$1", id)
	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.PasswordHash)

	return user, err
}

func (pg *postgres) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	row := pg.db.QueryRow(context.Background(), "select * from public.user where email=$1", email)
	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.PasswordHash)

	return user, err
}
