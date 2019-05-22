package api

import . "tunaiku-api/struct"

type RepositoryLoaning interface {
	Get() ([]*Loaning,  error)
	GetId(id uint) (*Loaning, error)
	Create(loaning *Loaning) error
	Update(loaning *Loaning) error
	Delete(loaning *Loaning) error
}

type RepositoryKtpDetail interface{
	Create(ktp uint) error
}
type RepositoryDate interface {
	Create(date *Date) error
	Update(date *Date) error
	Delete(date *Date) error
	Get() ([]*Date, error)
	GetId(id uint) (*Date, error)
}
type RepositoryMonth interface {
	Create(month *Month) error
	Update(month *Month) error
	Delete(month *Month) error
	Get() ([]*Month, error)
	GetId(id uint) (*Month, error)
}
type RepositoryProvincial interface {
	Create(provincial *Provincial) error
	Update(provincial *Provincial) error
	Delete(provincial *Provincial) error
	Get() ([]*Provincial, error)
	GetId(id uint) (*Provincial, error)
}
type RepositorySubDistrict interface {
	Create(subdistrict *SubDistrict) error
	Update(subdistrict *SubDistrict) error
	Delete(subdistrict *SubDistrict) error
	Get() ([]*SubDistrict, error)
	GetId(id uint) (*SubDistrict, error)
}
type RepositoryYear interface {
	Create(year *Year) error
	Update(year *Year) error
	Delete(year *Year) error
	Get() ([]*Year, error)
	GetId(id uint) (*Year, error)
}