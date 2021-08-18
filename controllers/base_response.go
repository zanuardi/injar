package controller

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Meta struct {
		Status   int      `json:"code"`
		Message  string   `json:"message"`
		Messages []string `json:"messages,omitempty"`
	} `json:"meta"`
	Data       interface{} `json:"data"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

func NewSuccessResponse(c echo.Context, param interface{}) error {
	response := BaseResponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "Success"
	response.Data = param

	return c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Meta.Status = status
	response.Meta.Message = "Something not right"
	response.Meta.Messages = []string{err.Error()}

	return c.JSON(http.StatusOK, response)
}

type Pagination struct {
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
}

func PaginationRes(page, total, limit int) Pagination {
	lastPage := total / limit
	if total%limit > 0 {
		lastPage = lastPage + 1
	}

	pagination := Pagination{
		CurrentPage: page,
		LastPage:    lastPage,
		Total:       total,
		PerPage:     limit,
	}
	return pagination
}

func NewSuccessResponsePagination(c echo.Context, data interface{}, pagination Pagination) error {
	response := BaseResponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "Success"
	response.Data = data
	response.Pagination = &pagination

	return c.JSON(http.StatusOK, response)
}
