package workbook_categories

import (
	"errors"
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type IndexActionCommand struct {
	WorkbookId uuid.UUID
}

type IndexAction struct {
	WorkbookCategoryRepository repositories.WorkbookCategoryRepository
}

func (a *IndexAction) Execute(command *IndexActionCommand) ([]*WorkbookCategoryDto, usecase_error.UsecaseErrorGroup) {
	workbookCategories := a.WorkbookCategoryRepository.FindByWorkbookId(command.WorkbookId)

	if len(workbookCategories) == 0 {
		return nil, usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, errors.New("workbookCategories not found")))
	}

	categoryDtos := lo.Map(workbookCategories, func(root *entities.WorkbookCategory, _ int) *WorkbookCategoryDto {
		children1Dtos := lo.Map(root.Children(), func(children1 *entities.WorkbookCategory, _ int) *WorkbookCategoryDto {
			children2Dtos := lo.Map(children1.Children(), func(children2 *entities.WorkbookCategory, _ int) *WorkbookCategoryDto {
				children3Dtos := lo.Map(children2.Children(), func(children3 *entities.WorkbookCategory, _ int) *WorkbookCategoryDto {
					return &WorkbookCategoryDto{
						WorkbookCategoryId: children3.Id(),
						Children:           make([]*WorkbookCategoryDto, 0),
						Name:               children3.Name(),
					}
				})

				return &WorkbookCategoryDto{
					WorkbookCategoryId: children2.Id(),
					Children:           children3Dtos,
					Name:               children2.Name(),
				}
			})

			return &WorkbookCategoryDto{
				WorkbookCategoryId: children1.Id(),
				Children:           children2Dtos,
				Name:               children1.Name(),
			}
		})

		return &WorkbookCategoryDto{
			WorkbookCategoryId: root.Id(),
			Children:           children1Dtos,
			Name:               root.Name(),
		}
	})

	return categoryDtos, nil
}
