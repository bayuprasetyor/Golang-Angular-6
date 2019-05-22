package usecase

import (
	"tunaiku-api/api"
	"tunaiku-api/struct"
)


type DateUseCase struct {
	repoDate api.RepositoryDate
}

func (du *DateUseCase) CreateDate(date *_struct.Date) error {
	return du.repoDate.Create(date)
}

func (du *DateUseCase) UpdateDate(date *_struct.Date) error {
	return du.repoDate.Update(date)
}

func (du *DateUseCase) DeleteDate(date *_struct.Date) error {
	return du.repoDate.Delete(date)
}

func (du *DateUseCase) GetDates() ([]*_struct.Date, error) {
	res, err := du.repoDate.Get()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (du *DateUseCase) GetDatesId(id uint) (*_struct.Date, error) {
	res, err := du.repoDate.GetId(id)

	if err != nil {
		return nil, err
	}

	return res, err
}

func NewDateUseCase(repoDate api.RepositoryDate) api.DateUseCase{
	return &DateUseCase{
		repoDate,
	}
}