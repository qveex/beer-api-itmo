package domain

func (s AuthService) GetUserId() (int64, error) {
	return s.repo.GetUserId()
}
