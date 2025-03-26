package problems

import "study-pal-backend/app/usecases/shared/usecase_error"

type problem struct {
	Name string
}
type workbookCategoryDetail struct {
	Name     string
	Problems []*problem
}

type workbookCategory struct {
	Name                    string
	Problems                []*problem
	WorkbookCategoryDetails []*workbookCategoryDetail
}

type CreateActionCommand struct {
	Problems           []*problem
	WorkbookCategories []*workbookCategory
}

type CreateAction struct {
}

func (a *CreateAction) Execute(command *CreateActionCommand) usecase_error.UsecaseErrorGroup {
	return nil
}
