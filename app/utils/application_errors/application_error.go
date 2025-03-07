package application_errors

type applicationErrorType int

type ApplicationError interface {
	Kind() applicationErrorType
	Error() string
}

const (
	ClientInputValidation applicationErrorType = iota
	DatabaseConnection
	DataNotFound
	FatalError
)

type applicationError struct {
	kind    applicationErrorType
	message string
}

func NewApplicationError(kind applicationErrorType, err error) *applicationError {
	return &applicationError{
		kind:    kind,
		message: err.Error(),
	}
}

func NewClientInputValidationApplicationError(err error) *applicationError {
	return &applicationError{
		kind:    ClientInputValidation,
		message: err.Error(),
	}
}

func NewDataNotFoundApplicationError(err error) *applicationError {
	return &applicationError{
		kind:    DataNotFound,
		message: err.Error(),
	}
}

func NewDatabaseConnectionApplicationError(err error) *applicationError {
	return &applicationError{
		kind:    DatabaseConnection,
		message: err.Error(),
	}
}

func NewFatalApplicationError(err error) *applicationError {
	return &applicationError{
		kind:    FatalError,
		message: err.Error(),
	}
}

func (apperr *applicationError) Kind() applicationErrorType {
	return apperr.kind
}

func (apperr *applicationError) Error() string {
	return apperr.message
}
