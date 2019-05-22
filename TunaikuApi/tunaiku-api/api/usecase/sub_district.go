package usecase

import (
	"tunaiku-api/api"
	"tunaiku-api/struct"
)

type SubDistrictUseCase struct {
	repoSubDistrict api.RepositorySubDistrict
}

func (su *SubDistrictUseCase) CreateSubDistrict(subdistrict *_struct.SubDistrict) error {
	return su.repoSubDistrict.Create(subdistrict)
}

func (su *SubDistrictUseCase) UpdateSubDistrict(subdistrict *_struct.SubDistrict) error {
	return su.repoSubDistrict.Update(subdistrict)
}

func (su *SubDistrictUseCase) DeleteSubDistrict(subdistrict *_struct.SubDistrict) error {
	return su.repoSubDistrict.Delete(subdistrict)
}

func (su *SubDistrictUseCase) GetSubDistricts() ([]*_struct.SubDistrict, error) {
	res, err := su.repoSubDistrict.Get()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (su *SubDistrictUseCase) GetSubDistrictsId(id uint) (*_struct.SubDistrict, error) {
	res, err := su.repoSubDistrict.GetId(id)

	if err != nil {
		return nil, err
	}

	return res, err
}

func NewSubDistrictUseCase(repoSubDistrict api.RepositorySubDistrict) api.SubDistrictUseCase {
	return &SubDistrictUseCase{
		repoSubDistrict,
	}
}
