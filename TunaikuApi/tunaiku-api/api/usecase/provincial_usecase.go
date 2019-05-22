package usecase

import (
	"tunaiku-api/api"
	"tunaiku-api/struct"
)

type ProvincialUseCase struct {
	repoProvincial api.RepositoryProvincial
}

func (pu *ProvincialUseCase) CreateProvincial(provincial *_struct.Provincial) error {
	return pu.repoProvincial.Create(provincial)
}

func (pu *ProvincialUseCase) UpdateProvincial(provincial *_struct.Provincial) error {
	return pu.repoProvincial.Update(provincial)
}

func (pu *ProvincialUseCase) DeleteProvincial(provincial *_struct.Provincial) error {
	return pu.repoProvincial.Delete(provincial)
}

func (pu *ProvincialUseCase) GetProvincials() ([]*_struct.Provincial, error) {
	res, err := pu.repoProvincial.Get()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *ProvincialUseCase) GetProvincialsId(id uint) (*_struct.Provincial, error) {
	res, err := pu.repoProvincial.GetId(id)

	if err != nil {
		return nil, err
	}

	return res, err
}

func NewProvincialUseCase(repoProvincial api.RepositoryProvincial) api.ProvincialUseCase{
	return &ProvincialUseCase{
		repoProvincial,
	}
}
