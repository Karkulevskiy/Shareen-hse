package utils

import (
	"net/http"
	"shareen/src/internal/models"

	"github.com/google/uuid"
)

func ValidateId(id string) *models.ResponseError {
	if id == "" {
		return &models.ResponseError{
			Message: "id can't be null",
			Status:  http.StatusBadRequest,
		}
	}
	err := uuid.Validate(id)
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}
