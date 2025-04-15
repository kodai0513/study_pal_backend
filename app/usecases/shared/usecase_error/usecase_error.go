package usecase_error

type usecaseErrorType int

type UsecaseError interface {
	UsecaseErrorType() usecaseErrorType
	Error() string
}

const (
	// 無効なパラメータ
	InvalidParameter usecaseErrorType = iota
	// データが見つからない
	QueryDataNotFoundError
	// 許可されていない操作
	UnPermittedOperation
	// データベースエラー
	DatabaseError
)

type usecaseError struct {
	usecaseErrorType usecaseErrorType
	message          string
}

func NewUsecaseError(usecaseErrorType usecaseErrorType, err error) UsecaseError {
	return &usecaseError{
		usecaseErrorType: usecaseErrorType,
		message:          err.Error(),
	}
}

func (u *usecaseError) UsecaseErrorType() usecaseErrorType {
	return u.usecaseErrorType
}

func (u *usecaseError) Error() string {
	return u.message
}
