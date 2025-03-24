package problems

import "study-pal-backend/app/usecases/shared/usecase_error"

type problem struct {
	Name string
}
type workbookCategoryClassification struct {
	Name     string
	Problems []*problem
}

type workbookCategory struct {
	Name                            string
	Problems                        []*problem
	WorkbookCategoryClassifications []*workbookCategoryClassification
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
