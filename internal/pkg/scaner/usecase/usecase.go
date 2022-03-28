package usecase

import (
	"http-proxy/internal/models"
	"http-proxy/internal/pkg/scaner"
)

type UseCase struct {
	repo scaner.Repository
}

func NewRepoUsecase(repo scaner.Repository) scaner.UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) GetRequests() ([]models.Request, models.StatusCode) {
	return uc.repo.GetRequests()
}
