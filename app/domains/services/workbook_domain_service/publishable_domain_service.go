package workbook_domain_services

import (
	"errors"
	"study-pal-backend/app/domains/repositories"

	"github.com/google/uuid"
)

type PublishableDomainService struct {
	problemRespository repositories.ProblemRepository
}

func NewPublishableDomainService(problemRespository repositories.ProblemRepository) *PublishableDomainService {
	return &PublishableDomainService{
		problemRespository: problemRespository,
	}
}

func (u *PublishableDomainService) Execute(workbookId uuid.UUID) error {
	exist := u.problemRespository.ExistByWorkbookId(workbookId)
	if exist {
		return nil
	}

	return errors.New("this workbook cannot be published")
}
