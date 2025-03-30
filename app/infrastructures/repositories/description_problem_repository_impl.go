package repositories

import (
	"context"
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/ent"

	"github.com/google/uuid"
)

type DescriptionProblemRepositoryImpl struct {
	Client *ent.Client
	Ctx    context.Context
}

func (p *DescriptionProblemRepositoryImpl) CreateBulk(problems []*entities.DescriptionProblem) {

}

func (p *DescriptionProblemRepositoryImpl) ExistByWorkbookId(workbookId uuid.UUID) bool {
	return false
}
