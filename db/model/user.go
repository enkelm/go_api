package model

import "github.com/jackc/pgtype"

type User struct {
	Id           pgtype.UUID
	FirstName    pgtype.Varchar
	LastName     pgtype.Varchar
	Email        pgtype.Varchar
	PasswordHash pgtype.Varchar
}

type UserDto struct {
	FirstName string
	LastName  string
	Email     string
}

func (u *User) Data() UserDto {
	return UserDto{FirstName: u.FirstName.String, LastName: u.LastName.String, Email: u.Email.String}
}
