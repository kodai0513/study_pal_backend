package mappers

import (
	"fmt"
	"net/http"
	"study-pal-backend/app/usecases/shared/usecase_errors"
)

func UsecaseErrorToHttpStatus(usecaseErrGroup usecase_errors.UsecaseErrorGroup) int {
	switch usecaseErrGroup.UsecaseErrorType() {
	case usecase_errors.InvalidParameter:
		return http.StatusBadRequest
	case usecase_errors.QueryDataNotFoundError:
		return http.StatusNotFound
	case usecase_errors.UnPermittedOperation:
		return http.StatusUnauthorized
	default:
		panic(fmt.Sprintf("unexpected error type: %v in AppErrorToHttpStatus", usecaseErrGroup.UsecaseErrorType()))
	}
}
