package http

import (
	"github.com/labstack/echo"
	"tunaiku-api/api"
)

type KtpDetailHandler struct {
	KtpDetailUseCase api.KtpDetailUseCase
}

func NewKtpDetailHanlder(e *echo.Echo, ktpdetailusecase api.KtpDetailUseCase) {
	handler := &KtpDetailHandler{KtpDetailUseCase:ktpdetailusecase}

	e.POST("/ktp-detail", handler.CreateKtpDetail)
}

func (kh *KtpDetailHandler) CreateKtpDetail (c echo.Context) error {
	return nil
}
