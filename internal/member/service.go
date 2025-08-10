package member

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(request *CreateMemberRequest) error {
	return s.repo.Save(request)
}

func (s *Service) GetAll() ([]*GetMemberDTO, error) {
	return s.repo.FindAll()
}

func (s *Service) Update(id string, request *UpdateMemberRequest) error {
	return s.repo.Update(id, request)
}

func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}
