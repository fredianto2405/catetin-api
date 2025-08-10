package auth

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindByEmail(email string) (*UserDTO, error) {
	sql := `select id,
			email,
			password,
			role,
			failed_login_attempts,
			is_locked 
		from catetin.users
		where email = $1`

	var user UserDTO
	err := r.db.Get(&user, sql, email)
	return &user, err
}

func (r *Repository) UpdateFailedLoginAttempts(email string, isLocked bool) error {
	sql := `update catetin.users
		set failed_login_attempts = failed_login_attempts + 1,
			is_locked = $1
		where email = $2`

	_, err := r.db.Exec(sql, isLocked, email)
	return err
}

func (r *Repository) ResetFailedLoginAttempts(email string) error {
	sql := `update catetin.users
		set failed_login_attempts = 0,
			is_locked = false
		where email = $1
		and failed_login_attempts > 0`

	_, err := r.db.Exec(sql, email)
	return err
}

func (r *Repository) UpdatePassword(email string, password string) error {
	sql := `update catetin.users
		set password = $1,
		    updated_at = current_timestamp
		where email = $2`

	_, err := r.db.Exec(sql, password, email)
	return err
}
