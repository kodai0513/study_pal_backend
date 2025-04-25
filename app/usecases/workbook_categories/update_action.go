package workbook_categories

import (
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/models/value_objects/workbook_categories"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/permission_guard"
	"study-pal-backend/app/usecases/shared/trancaction"
	"study-pal-backend/app/usecases/shared/usecase_error"
	"study-pal-backend/app/utils/type_converts"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type WorkbookCategory struct {
	WorkbookCategoryId *uuid.UUID
	Children           []*WorkbookCategory
	Name               string
}

type UpdateActionCommand struct {
	UserId             uuid.UUID
	WorkbookId         uuid.UUID
	WorkbookCategories []*WorkbookCategory
}

type UpdateAction struct {
	PermissionGuard            permission_guard.WorkbookPermissionGuard
	Tx                         trancaction.Tx
	WorkbookCategoryRepository repositories.WorkbookCategoryRepository
}

func (a *UpdateAction) Execute(command *UpdateActionCommand) ([]*WorkbookCategoryDto, usecase_error.UsecaseErrorGroup) {
	err := a.PermissionGuard.Check("update:workbook-categories", command.UserId, command.WorkbookId)
	if err != nil {
		return nil, usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.UnPermittedOperation, err))
	}
	invalidUsecaseErrGroup := usecase_error.NewUsecaseErrorGroup(usecase_error.InvalidParameter)

	categories := lo.Map(command.WorkbookCategories, func(root *WorkbookCategory, _ int) *entities.WorkbookCategory {
		children1Categories := lo.Map(root.Children, func(children1 *WorkbookCategory, _ int) *entities.WorkbookCategory {
			children1Name, err := workbook_categories.NewName(children1.Name)
			if err != nil {
				invalidUsecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
			}
			children2Categories := lo.Map(children1.Children, func(children2 *WorkbookCategory, _ int) *entities.WorkbookCategory {
				children2Name, err := workbook_categories.NewName(children2.Name)
				if err != nil {
					invalidUsecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
				}
				children3Categories := lo.Map(children2.Children, func(children3 *WorkbookCategory, _ int) *entities.WorkbookCategory {
					children3Name, err := workbook_categories.NewName(children3.Name)
					if err != nil {
						invalidUsecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
					}
					return entities.NewWorkbookCategory(
						type_converts.PointerUuidToUuidIfNilWhenNewUuid(children3.WorkbookCategoryId),
						make([]*entities.WorkbookCategory, 0),
						children3Name,
						command.WorkbookId,
					)
				})

				return entities.NewWorkbookCategory(
					type_converts.PointerUuidToUuidIfNilWhenNewUuid(children2.WorkbookCategoryId),
					children3Categories,
					children2Name,
					command.WorkbookId,
				)
			})

			return entities.NewWorkbookCategory(
				type_converts.PointerUuidToUuidIfNilWhenNewUuid(children1.WorkbookCategoryId),
				children2Categories,
				children1Name,
				command.WorkbookId,
			)
		})

		rootName, err := workbook_categories.NewName(root.Name)
		if err != nil {
			invalidUsecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
		}
		return entities.NewWorkbookCategory(
			type_converts.PointerUuidToUuidIfNilWhenNewUuid(root.WorkbookCategoryId),
			children1Categories,
			rootName,
			command.WorkbookId,
		)
	})

	if invalidUsecaseErrGroup.IsError() {
		return nil, invalidUsecaseErrGroup
	}

	results := []*entities.WorkbookCategory{}
	err = trancaction.WithTx(a.Tx, func() {
		results = a.WorkbookCategoryRepository.UpsertAndDeleteBulk(categories, command.WorkbookId)
	})

	if err != nil {
		return nil, usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.DatabaseError, err))
	}

	categoryDtos := lo.Map(results, func(root *entities.WorkbookCategory, _ int) *WorkbookCategoryDto {
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
