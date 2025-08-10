package member

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(request *CreateMemberRequest) error {
	sql := `insert into catetin.members(name, role, email, phone, notes)
		values(:name, :role, :email, :phone, :notes)`

	_, err := r.db.NamedExec(sql, request)
	return err
}

func (r *Repository) FindAll() ([]*GetMemberDTO, error) {
	sql := `select id, 
			name,
			role,
			email,
			phone,
			notes
		from catetin.members
		order by name asc`

	var members []*GetMemberDTO
	err := r.db.Select(&members, sql)
	return members, err
}

func (r *Repository) Update(id string, request *UpdateMemberRequest) error {
	sql := `update catetin.members
		set name = $1,
			role = $2,
			email = $3,
			phone = $4,
			notes = $5,
			updated_at = current_timestamp
		where id = $6`

	_, err := r.db.Exec(sql, request.Name, request.Role, request.Email, request.Phone, request.Notes, id)
	return err
}

func (r *Repository) Delete(id string) error {
	sql := `delete from catetin.members where id = $1`
	_, err := r.db.Exec(sql, id)
	return err
}
