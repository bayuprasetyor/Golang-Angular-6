package usecase

import (
	"tunaiku-api/api"
	"tunaiku-api/struct"
)

type YearUseCase struct {
	repoYear api.RepositoryYear
}

func (yu *YearUseCase) CreateYear(year *_struct.Year) error {
	return yu.repoYear.Create(year)
}

func (yu *YearUseCase) UpdateYear(year *_struct.Year) error {
	return yu.repoYear.Update(year)
}

func (yu *YearUseCase) DeleteYear(year *_struct.Year) error {
	return yu.repoYear.Delete(year)
}

func (yu *YearUseCase) GetYears() ([]*_struct.Year, error) {
	res, err := yu.repoYear.Get()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (yu *YearUseCase) GetYearsId(id uint) (*_struct.Year, error) {
	res, err := yu.repoYear.GetId(id)

	if err != nil {
		return nil, err
	}

	return res, err
}

func NewYearUseCase(repoYear api.RepositoryYear) api.YearUseCase {
	return &YearUseCase{
		repoYear,
	}
}

