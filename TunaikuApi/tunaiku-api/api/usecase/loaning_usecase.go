package usecase

import (
	"tunaiku-api/struct"
	"tunaiku-api/api"
)

type LoaningUseCase struct {
	repoLoaning api.RepositoryLoaning
}

func (au *LoaningUseCase) GetLoaningsId(id uint) (*_struct.Loaning, error) {
	res, err := au.repoLoaning.GetId(id)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (au *LoaningUseCase) CreateLoaning(loaning *_struct.Loaning) error {
	return au.repoLoaning.Create(loaning)
}

func (au *LoaningUseCase) DeleteLoaning(loaning *_struct.Loaning) error {
	return au.repoLoaning.Delete(loaning)
}

func (au *LoaningUseCase) GetLoanings() ([]*_struct.Loaning, error) {
	res, err := au.repoLoaning.Get()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (au *LoaningUseCase) UpdateLoaning(loaning *_struct.Loaning) error {
	return au.repoLoaning.Update(loaning)
}

func NewLoaningUseCase(repoLoaning api.RepositoryLoaning) api.LoaningUseCase {
	return &LoaningUseCase{
		repoLoaning,
	}
}