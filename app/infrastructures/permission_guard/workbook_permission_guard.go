package permission_guard

import (
	"errors"
	"study-pal-backend/app/usecases/shared/permission_guard"
	"study-pal-backend/ent"
	"study-pal-backend/ent/permission"
	"study-pal-backend/ent/workbook"
	"study-pal-backend/ent/workbookmember"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type workbookPermissionGuard struct {
	client *ent.Client
	ctx    *gin.Context
}

func NewWorkbookPermissionGuard(client *ent.Client, ctx *gin.Context) permission_guard.WorkbookPermissionGuard {
	return &workbookPermissionGuard{
		client: client,
		ctx:    ctx,
	}
}

func (w *workbookPermissionGuard) Check(permissionSource string, userId uuid.UUID, workbookId uuid.UUID) error {
	result := w.client.Workbook.Query().
		Where(workbook.IDEQ(workbookId)).
		WithWorkbookMembers(func(wmq *ent.WorkbookMemberQuery) {
			wmq.Where(workbookmember.UserIDEQ(userId)).
				WithRole(func(rq *ent.RoleQuery) {
					rq.WithPermissions(func(pq *ent.PermissionQuery) {
						pq.Where(permission.NameEQ(permissionSource))
					})
				})
		}).
		FirstX(w.ctx)

	if result == nil || len(result.Edges.WorkbookMembers) == 0 || len(result.Edges.WorkbookMembers[0].Edges.Role.Edges.Permissions) == 0 {
		return errors.New("you do not have the necessary authorization")
	}

	return nil
}
