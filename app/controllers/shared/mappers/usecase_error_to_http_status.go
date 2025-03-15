package mappers

import (
	"fmt"
	"net/http"
	"study-pal-backend/app/usecases/shared/usecase_error"
)

func UsecaseErrorToHttpStatus(usecaseErrGroup usecase_error.UsecaseErrorGroup) int {
	switch usecaseErrGroup.UsecaseErrorType() {
	case usecase_error.InvalidParameter:
		return http.StatusBadRequest
	case usecase_error.QueryDataNotFoundError:
		return http.StatusNotFound
	case usecase_error.UnPermittedOperation:
		return http.StatusUnauthorized
	default:
		panic(fmt.Sprintf("unexpected error type: %v in AppErrorToHttpStatus", usecaseErrGroup.UsecaseErrorType()))
	}
}
