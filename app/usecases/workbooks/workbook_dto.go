package workbooks

import "github.com/google/uuid"

type WorkbookDto struct {
	Id          uuid.UUID
	Description string
	IsPublic    bool
	Title       string
	UserId      uuid.UUID
}
