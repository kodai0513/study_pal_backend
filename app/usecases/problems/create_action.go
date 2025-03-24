package problems

import "study-pal-backend/app/usecases/shared/usecase_error"

type problem struct {
	name string
}

func NewProblem(name string) *problem {
	return &problem{name: name}
}

type workbookCategoryClassification struct {
	name     string
	problems []*problem
}

func NewWorkbookCategoryClassification(name string, problems []*problem) *workbookCategoryClassification {
	return &workbookCategoryClassification{
		name:     name,
		problems: problems,
	}
}

type workbookCategory struct {
	name                            string
	problems                        []*problem
	workbookCategoryClassifications []*workbookCategoryClassification
}

func NewWorkbookCategory(name string, problems []*problem, workbookCategoryClassifications []*workbookCategoryClassification) *workbookCategory {
	return &workbookCategory{
		name:                            name,
		problems:                        problems,
		workbookCategoryClassifications: workbookCategoryClassifications,
	}
}

type CreateActionCommand struct {
	problems           []*problem
	workbookCategories []*workbookCategory
}

func NewCreateActionCommand(problems []*problem, workbookCategories []*workbookCategory) *CreateActionCommand {
	return &CreateActionCommand{
		problems:           problems,
		workbookCategories: workbookCategories,
	}
}

type CreateAction struct {
}

func NewCreateAction() *CreateAction {
	return &CreateAction{}
}

func (a *CreateAction) Execute(command *CreateActionCommand) usecase_error.UsecaseErrorGroup {
	return nil
}
