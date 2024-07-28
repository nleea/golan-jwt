package repository

import (
	"context"
	"database/sql"
	"fmt"
	"learn/jwt/domain"
)

type userRepository struct {
	db    *sql.DB
	table string
}

// Create implements domain.UserRepository.
func (u *userRepository) Create(c context.Context, user *domain.User) error {

	query := `INSERT into "%s" (name, email, password) VALUES ($1, $2, $3)`
	query = fmt.Sprintf(query, u.table)

	_, err := u.db.Exec(query, user.Name, user.Email, user.Password)

	return err

}

// Fetch implements domain.UserRepository.
func (u *userRepository) Fetch(c context.Context) ([]domain.User, error) {
	panic("unimplemented")
}

// GetByEmail implements domain.UserRepository.
func (u *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	user := &domain.User{}

	query := `SELECT * FROM "%s" WHERE email=$1`
	query = fmt.Sprintf(query, u.table)

	row := u.db.QueryRowContext(c, query, email)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return *user, fmt.Errorf("user not found")
		}
		return *user, err
	}

	return *user, nil
}

// GetByID implements domain.UserRepository.
func (u *userRepository) GetByID(c context.Context, id string) (domain.User, error) {
	user := &domain.User{}

	query := "SELECT * FROM %s WHERE email=%s"
	query = fmt.Sprint(query, u.table)
	u.db.Exec(query)

	return *user, nil
}

func NewUserRepository(db *sql.DB, table string) domain.UserRepository {
	return &userRepository{
		db:    db,
		table: table,
	}
}
