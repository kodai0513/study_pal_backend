package usecase_error

import (
	"github.com/samber/lo"
)

type UsecaseErrorGroup interface {
	AddOnlySameUsecaseError(usecaseErr UsecaseError)
	Errors() []string
	IsError() bool
	UsecaseErrorType() usecaseErrorType
}

type usecaseErrorGroup struct {
	usecaseErrorType usecaseErrorType
	errors           []UsecaseError
}

func NewUsecaseErrorGroup(usecaseErrorType usecaseErrorType) UsecaseErrorGroup {
	return &usecaseErrorGroup{
		usecaseErrorType: usecaseErrorType,
		errors:           []UsecaseError{},
	}
}

func NewUsecaseErrorGroupWithMessage(usecaseErr UsecaseError) UsecaseErrorGroup {
	if usecaseErr == nil {
		panic("do not pass nil as an argument")
	}

	return &usecaseErrorGroup{
		usecaseErrorType: usecaseErr.UsecaseErrorType(),
		errors:           []UsecaseError{usecaseErr},
	}
}

func (a *usecaseErrorGroup) AddOnlySameUsecaseError(usecaseErr UsecaseError) {
	if len(a.errors) > 0 && usecaseErr.UsecaseErrorType() != a.errors[0].UsecaseErrorType() {
		panic("do not add different error types")
	}

	a.errors = append(a.errors, usecaseErr)
}

func (a *usecaseErrorGroup) Errors() []string {
	return lo.Map(
		a.errors,
		func(err UsecaseError, index int) string {
			return err.Error()
		},
	)
}

func (a *usecaseErrorGroup) IsError() bool {
	return len(a.errors) > 0
}

func (a *usecaseErrorGroup) UsecaseErrorType() usecaseErrorType {
	return a.usecaseErrorType
}
