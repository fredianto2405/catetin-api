package auth

import (
	"errors"
	"github.com/fredianto2405/catetin-api/pkg/password"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Login(request *LoginRequest) (*UserDTO, error) {
	user, err := s.repo.FindByEmail(request.Email)
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	if user.FailedLoginAttempts >= 3 {
		return nil, errors.New("akun anda diblokir")
	}

	isPasswordMatch := password.CheckPasswordHash(request.Password, user.Password)
	if !isPasswordMatch {
		isLocked := (user.FailedLoginAttempts + 1) == 3
		if err = s.repo.UpdateFailedLoginAttempts(request.Email, isLocked); err != nil {
			return nil, err
		}
		return nil, errors.New("password tidak valid")
	}

	if err = s.repo.ResetFailedLoginAttempts(request.Email); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) ChangePassword(email string, request *ChangePasswordRequest) error {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return errors.New("user tidak ditemukan")
	}

	isPasswordMatch := password.CheckPasswordHash(request.OldPassword, user.Password)
	if !isPasswordMatch {
		return errors.New("password tidak valid")
	}

	if request.NewPassword != request.ConfirmNewPassword {
		return errors.New("konfirmasi password tidak valid")
	}

	if err = password.Validate(request.NewPassword); err != nil {
		return err
	}

	var hashedPassword string
	hashedPassword, err = password.HashPassword(request.NewPassword)
	if err != nil {
		return err
	}

	if err = s.repo.UpdatePassword(user.Email, hashedPassword); err != nil {
		return err
	}

	return nil
}
