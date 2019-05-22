package usecase

import (
	"tunaiku-api/api"
	"tunaiku-api/struct"
)

type MonthUseCase struct {
	repoMonth api.RepositoryMonth
}

func (mu *MonthUseCase) CreateMonth(month *_struct.Month) error {
	return mu.repoMonth.Create(month)
}

func (mu *MonthUseCase) UpdateMonth(month *_struct.Month) error {
	return mu.repoMonth.Update(month)
}

func (mu *MonthUseCase) DeleteMonth(month *_struct.Month) error {
	return mu.repoMonth.Delete(month)
}

func (mu *MonthUseCase) GetMonths() ([]*_struct.Month, error) {
	res, err := mu.repoMonth.Get()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (mu *MonthUseCase) GetMonthsId(id uint) (*_struct.Month, error) {
	res, err := mu.repoMonth.GetId(id)

	if err != nil {
		return nil, err
	}

	return res, err
}

func NewMonthUseCase(repoMonth api.RepositoryMonth) api.MonthUseCase {
	return &MonthUseCase{
		repoMonth,
	}
}