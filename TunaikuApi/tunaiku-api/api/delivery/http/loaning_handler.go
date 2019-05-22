package http

import (
	"github.com/labstack/echo"
	"strconv"
	. "tunaiku-api/api"
	. "tunaiku-api/struct"
	. "tunaiku-api/api/delivery/utils"
	"net/http"
)

type LoaningHandler struct {
	LoaningUseCase LoaningUseCase
}

func NewLoaningHanlder(e *echo.Echo, loaningusecase LoaningUseCase) {
	handler := &LoaningHandler{LoaningUseCase:loaningusecase}

	e.POST("/api/loanings", handler.CreateLoaning)
	e.GET("/api/loanings", handler.GetLoanings)
	e.GET("/api/loanings/:id", handler.GetLoaningsId)
	e.PUT("/api/loanings/:id", handler.GetLoaningUpdate)
	e.DELETE("/api/loanings/:id", handler.GetLoaningDelete)
}

func (lh *LoaningHandler) CreateLoaning(c echo.Context) error {
	var loaning Loaning

	err := c.Bind(&loaning)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := IsLoaningRequestValid(&loaning); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = lh.LoaningUseCase.CreateLoaning(&loaning)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, loaning)
}

func (lh *LoaningHandler) GetLoanings(c echo.Context) error {
	loanings, err := lh.LoaningUseCase.GetLoanings()

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, loanings)
}
func (lh *LoaningHandler) GetLoaningsId(c echo.Context) error {

	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	loanings, err := lh.LoaningUseCase.GetLoaningsId(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, loanings)
}
func (lh *LoaningHandler) GetLoaningUpdate(c echo.Context) error {
	var loaning_ Loaning

	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	_, err = lh.LoaningUseCase.GetLoaningsId(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	loaning_.ID = id

	err = c.Bind(&loaning_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := IsRequestForLoaningValid(&loaning_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = lh.LoaningUseCase.UpdateLoaning(&loaning_)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, loaning_)
}
func (lh *LoaningHandler) GetLoaningDelete(c echo.Context) error {
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	courier, err := lh.LoaningUseCase.GetLoaningsId(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	err = lh.LoaningUseCase.DeleteLoaning(courier)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(200, "Delete Success")
}
