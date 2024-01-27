package db

import (
	"context"
	"log"

	"github.com/enkelm/go_api/db/model"
	"github.com/enkelm/go_api/util"

	"github.com/jackc/pgx/v5"
)

func (pg *postgres) SignUp(ctx context.Context, u *model.UserDto, pass string) error {
	hashedPass, err := util.HashPassword(pass)
	if err != nil {
		log.Println("Unable to hash password: %w", err)
		return err
	}

	query := "INSERT INTO public.user (first_name, last_name, email, password_hash) VALUES (@firstName, @lastName, @email, @passwordHash)"
	args := pgx.NamedArgs{
		"firstName":    u.FirstName,
		"lastName":     u.LastName,
		"email":        u.Email,
		"passwordHash": hashedPass,
	}

	_, err = pg.db.Exec(ctx, query, args)
	if err != nil {
		log.Println("Unable to insert user: %w", err)
		return err
	}

	return nil
}

func (pg *postgres) SignIn(ctx context.Context, email, pass string) bool {
	user, err := pg.GetUserByEmail(email)
	if err != nil {
		log.Println("User does not exist: %w", err)
		return false
	}

	return util.CheckPasswordHash(pass, user.PasswordHash.String)
}
