package permission_guard

import "github.com/google/uuid"

type WorkbookPermissionGuard interface {
	Check(permissionResource string, userId uuid.UUID, workbookId uuid.UUID) error
}
