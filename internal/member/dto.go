package member

type CreateMemberRequest struct {
	Name  string `db:"name" json:"name" validate:"required"`
	Role  string `db:"role" json:"role" validate:"required"`
	Email string `db:"email" json:"email" validate:"required"`
	Phone string `db:"phone" json:"phone" validate:"required"`
	Notes string `db:"notes" json:"notes"`
}

type UpdateMemberRequest struct {
	Name  string `json:"name" validate:"required"`
	Role  string `json:"role" validate:"required"`
	Email string `json:"email" validate:"required"`
	Phone string `json:"phone" validate:"required"`
	Notes string `json:"notes"`
}

type GetMemberDTO struct {
	ID    string `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Role  string `db:"role" json:"role"`
	Email string `db:"email" json:"email"`
	Phone string `db:"phone" json:"phone"`
	Notes string `db:"notes" json:"notes"`
}
