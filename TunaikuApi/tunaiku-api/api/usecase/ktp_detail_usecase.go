package usecase

import (
	"tunaiku-api/api"
)

type KtpDetailUseCase struct {
	repoKtpDetail api.RepositoryKtpDetail
}

func (ku *KtpDetailUseCase) CreateKtpDetail(ktp uint) error {
	return ku.repoKtpDetail.Create(ktp)
}

func NewKtpDetailUseCase(repoKtpDetail api.RepositoryKtpDetail) api.KtpDetailUseCase {
	return &KtpDetailUseCase{
		repoKtpDetail,
	}
}