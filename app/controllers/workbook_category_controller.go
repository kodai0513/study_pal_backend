package controllers

import (
	"net/http"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/controllers/shared/mappers"
	"study-pal-backend/app/infrastructures/repositories"
	"study-pal-backend/app/usecases/shared/trancaction"
	"study-pal-backend/app/usecases/workbook_categories"
	"study-pal-backend/app/utils/type_converts"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type WorkbookCategoryController struct {
	AppData *app_types.AppData
}

type WorkbookCategory struct {
	WorkbookCategoryId string              `json:"workbook_category_id"`
	Children           []*WorkbookCategory `json:"children"`
	Name               string              `json:"name"`
}

type IndexWorkbookCategoryResponse struct {
	WorkbookCategories []*WorkbookCategory `json:"workbook_categories"`
}

// workbook-category godoc
//
//	@Summary	API
//	@Description
//	@Tags		workbook-category
//	@Accept		json
//	@Produce	json
//	@Param		workbook_id	path		string	true	"Workbook ID"
//	@Success	200			{object}	IndexWorkbookCategoryResponse
//	@Failure	400			{object}	app_types.ErrorResponse
//	@Failure	401			{object}	app_types.ErrorResponse
//	@Failure	404			{object}	app_types.ErrorResponse
//	@Failure	500			{object}	app_types.ErrorResponse
//	@Router		/{workbook_id}/workbook-categories [get]
func (a *WorkbookCategoryController) Index(c *gin.Context) {
	workbookId, err := uuid.Parse(c.Param("workbook_id"))
	if err != nil {
		c.SecureJSON(
			http.StatusBadRequest,
			&app_types.ErrorResponse{
				Errors: []string{err.Error()},
			},
		)
		c.Abort()
		return
	}

	tx, err := a.AppData.Client().Tx(c)
	if err != nil {
		panic(err)
	}
	action := &workbook_categories.IndexAction{
		WorkbookCategoryRepository: repositories.NewWorkbookCategoryRepositoryImpl(tx, c),
	}

	categoryDtos, usecaseErrGroup := action.Execute(
		&workbook_categories.IndexActionCommand{
			WorkbookId: workbookId,
		},
	)

	if usecaseErrGroup != nil && usecaseErrGroup.IsError() {
		c.SecureJSON(
			mappers.UsecaseErrorToHttpStatus(usecaseErrGroup),
			&app_types.ErrorResponse{
				Errors: usecaseErrGroup.Errors(),
			},
		)
		c.Abort()
		return
	}

	response := &UpdateWorkbookCategoryResponse{
		WorkbookCategories: make([]*WorkbookCategory, 0),
	}
	response.WorkbookCategories = lo.Map(categoryDtos, func(root *workbook_categories.WorkbookCategoryDto, _ int) *WorkbookCategory {
		children1Categories := lo.Map(root.Children, func(children1 *workbook_categories.WorkbookCategoryDto, _ int) *WorkbookCategory {
			children2Categories := lo.Map(children1.Children, func(children2 *workbook_categories.WorkbookCategoryDto, _ int) *WorkbookCategory {
				children3Categories := lo.Map(children2.Children, func(children3 *workbook_categories.WorkbookCategoryDto, _ int) *WorkbookCategory {
					return &WorkbookCategory{
						WorkbookCategoryId: children3.WorkbookCategoryId.String(),
						Children:           make([]*WorkbookCategory, 0),
						Name:               children3.Name,
					}
				})

				return &WorkbookCategory{
					WorkbookCategoryId: children2.WorkbookCategoryId.String(),
					Children:           children3Categories,
					Name:               children2.Name,
				}
			})

			return &WorkbookCategory{
				WorkbookCategoryId: children1.WorkbookCategoryId.String(),
				Children:           children2Categories,
				Name:               children1.Name,
			}
		})

		return &WorkbookCategory{
			WorkbookCategoryId: root.WorkbookCategoryId.String(),
			Children:           children1Categories,
			Name:               root.Name,
		}
	})

	c.SecureJSON(
		http.StatusCreated,
		response,
	)

}

type UpdateWorkbookCategoryRequest struct {
	WorkbookCategories []*WorkbookCategory `json:"workbook_categories"`
}

type UpdateWorkbookCategoryResponse struct {
	WorkbookCategories []*WorkbookCategory `json:"workbook_categories"`
}

// workbook-category godoc
//
//	@Summary	API
//	@Description
//	@Tags		workbook-category
//	@Accept		json
//	@Produce	json
//	@Param		workbook_id	path		string							true	"Workbook ID"
//	@Param		request		body		UpdateWorkbookCategoryRequest	true	"カテゴリー更新リクエスト"
//	@Success	201			{object}	UpdateWorkbookCategoryResponse
//	@Failure	400			{object}	app_types.ErrorResponse
//	@Failure	401			{object}	app_types.ErrorResponse
//	@Failure	404			{object}	app_types.ErrorResponse
//	@Failure	500			{object}	app_types.ErrorResponse
//	@Router		/{workbook_id}/workbook-categories [put]
func (a *WorkbookCategoryController) Update(c *gin.Context) {
	var request UpdateWorkbookCategoryRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.SecureJSON(
			http.StatusBadRequest,
			&app_types.ErrorResponse{
				Errors: []string{err.Error()},
			},
		)
		c.Abort()
		return
	}
	workbookId, err := uuid.Parse(c.Param("workbook_id"))
	if err != nil {
		c.SecureJSON(
			http.StatusBadRequest,
			&app_types.ErrorResponse{
				Errors: []string{err.Error()},
			},
		)
		c.Abort()
		return
	}

	tx, err := a.AppData.Client().Tx(c)
	if err != nil {
		panic(err)
	}
	action := &workbook_categories.UpdateAction{
		Tx:                         trancaction.NewTx(tx),
		WorkbookCategoryRepository: repositories.NewWorkbookCategoryRepositoryImpl(tx, c),
	}
	command := &workbook_categories.UpdateActionCommand{
		WorkbookId:         workbookId,
		WorkbookCategories: make([]*workbook_categories.WorkbookCategory, 0),
	}
	invalidUuidErrors := make([]string, 0)
	command.WorkbookCategories = lo.Map(request.WorkbookCategories, func(root *WorkbookCategory, _ int) *workbook_categories.WorkbookCategory {
		children1Categories := lo.Map(root.Children, func(children1 *WorkbookCategory, _ int) *workbook_categories.WorkbookCategory {
			children2Categories := lo.Map(children1.Children, func(children2 *WorkbookCategory, _ int) *workbook_categories.WorkbookCategory {
				children3Categories := lo.Map(children2.Children, func(children3 *WorkbookCategory, _ int) *workbook_categories.WorkbookCategory {
					workbookCategoryUuid, err := type_converts.StringToUuidOrNil(children3.WorkbookCategoryId)
					if err != nil {
						invalidUuidErrors = append(invalidUuidErrors, err.Error())
					}
					return &workbook_categories.WorkbookCategory{
						WorkbookCategoryId: workbookCategoryUuid,
						Children:           make([]*workbook_categories.WorkbookCategory, 0),
						Name:               children3.Name,
					}
				})
				workbookCategoryUuid, err := type_converts.StringToUuidOrNil(children2.WorkbookCategoryId)
				if err != nil {
					invalidUuidErrors = append(invalidUuidErrors, err.Error())
				}
				return &workbook_categories.WorkbookCategory{
					WorkbookCategoryId: workbookCategoryUuid,
					Children:           children3Categories,
					Name:               children2.Name,
				}
			})
			workbookCategoryUuid, err := type_converts.StringToUuidOrNil(children1.WorkbookCategoryId)
			if err != nil {
				invalidUuidErrors = append(invalidUuidErrors, err.Error())
			}
			return &workbook_categories.WorkbookCategory{
				WorkbookCategoryId: workbookCategoryUuid,
				Children:           children2Categories,
				Name:               children1.Name,
			}
		})
		workbookCategoryUuid, err := type_converts.StringToUuidOrNil(root.WorkbookCategoryId)
		if err != nil {
			invalidUuidErrors = append(invalidUuidErrors, err.Error())
		}
		return &workbook_categories.WorkbookCategory{
			WorkbookCategoryId: workbookCategoryUuid,
			Children:           children1Categories,
			Name:               root.Name,
		}
	})

	if len(invalidUuidErrors) != 0 {
		c.SecureJSON(
			http.StatusBadRequest,
			&app_types.ErrorResponse{
				Errors: invalidUuidErrors,
			},
		)
		c.Abort()
		return
	}

	categoryDtos, usecaseErrGroup := action.Execute(command)

	if usecaseErrGroup != nil && usecaseErrGroup.IsError() {
		c.SecureJSON(
			mappers.UsecaseErrorToHttpStatus(usecaseErrGroup),
			&app_types.ErrorResponse{
				Errors: usecaseErrGroup.Errors(),
			},
		)
		c.Abort()
		return
	}

	response := &UpdateWorkbookCategoryResponse{
		WorkbookCategories: make([]*WorkbookCategory, 0),
	}
	response.WorkbookCategories = lo.Map(categoryDtos, func(root *workbook_categories.WorkbookCategoryDto, _ int) *WorkbookCategory {
		children1Categories := lo.Map(root.Children, func(children1 *workbook_categories.WorkbookCategoryDto, _ int) *WorkbookCategory {
			children2Categories := lo.Map(children1.Children, func(children2 *workbook_categories.WorkbookCategoryDto, _ int) *WorkbookCategory {
				children3Categories := lo.Map(children2.Children, func(children3 *workbook_categories.WorkbookCategoryDto, _ int) *WorkbookCategory {
					return &WorkbookCategory{
						WorkbookCategoryId: children3.WorkbookCategoryId.String(),
						Children:           make([]*WorkbookCategory, 0),
						Name:               children3.Name,
					}
				})

				return &WorkbookCategory{
					WorkbookCategoryId: children2.WorkbookCategoryId.String(),
					Children:           children3Categories,
					Name:               children2.Name,
				}
			})

			return &WorkbookCategory{
				WorkbookCategoryId: children1.WorkbookCategoryId.String(),
				Children:           children2Categories,
				Name:               children1.Name,
			}
		})

		return &WorkbookCategory{
			WorkbookCategoryId: root.WorkbookCategoryId.String(),
			Children:           children1Categories,
			Name:               root.Name,
		}
	})

	c.SecureJSON(
		http.StatusCreated,
		response,
	)
}
