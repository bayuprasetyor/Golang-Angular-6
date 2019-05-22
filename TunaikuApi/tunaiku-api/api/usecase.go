package api

import . "tunaiku-api/struct"

type LoaningUseCase interface {
	GetLoanings() ([]*Loaning, error)
	GetLoaningsId(id uint) (*Loaning, error)
	CreateLoaning(loaning *Loaning) error
	UpdateLoaning(loaning *Loaning) error
	DeleteLoaning(loaning *Loaning) error
}
type KtpDetailUseCase interface {
	CreateKtpDetail(ktp uint) error
}
type DateUseCase interface {
	CreateDate(date *Date) error
	UpdateDate(date *Date) error
	DeleteDate(date *Date) error
	GetDates() ([]*Date, error)
	GetDatesId(id uint) (*Date, error)
}
type MonthUseCase interface {
	CreateMonth(month *Month) error
	UpdateMonth(month *Month) error
	DeleteMonth(month *Month) error
	GetMonths() ([]*Month, error)
	GetMonthsId(id uint) (*Month, error)
}
type ProvincialUseCase interface {
	CreateProvincial(provincial *Provincial) error
	UpdateProvincial(provincial *Provincial) error
	DeleteProvincial(provincial *Provincial) error
	GetProvincials() ([]*Provincial, error)
	GetProvincialsId(id uint) (*Provincial, error)
}
type SubDistrictUseCase interface {
	CreateSubDistrict(subdistrict *SubDistrict) error
	UpdateSubDistrict(subdistrict *SubDistrict) error
	DeleteSubDistrict(subdistrict *SubDistrict) error
	GetSubDistricts() ([]*SubDistrict, error)
	GetSubDistrictsId(id uint) (*SubDistrict, error)
}
type YearUseCase interface {
	CreateYear(year *Year) error
	UpdateYear(year *Year) error
	DeleteYear(year *Year) error
	GetYears() ([]*Year, error)
	GetYearsId(id uint) (*Year, error)
}