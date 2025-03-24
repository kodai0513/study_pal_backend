package controllers

import (
	"study-pal-backend/app/app_types"

	"github.com/gin-gonic/gin"
)

type ProblemController struct {
	AppData *app_types.AppData
}

type Problem struct {
	Name string `json:"name"`
}

type WorkbookCategoryClassification struct {
	Name     string     `json:"name"`
	Problems []*Problem `json:"problems"`
}

type WorkbookCategory struct {
	Name                            string                            `json:"name"`
	Problems                        []*Problem                        `json:"problems"`
	WorkbookCategoryClassifications []*WorkbookCategoryClassification `json:"workbook_category_classifications"`
}

type ProblemCreateRequest struct {
	Problems           []*Problem          `json:"problems"`
	WorkbookCategories []*WorkbookCategory `json:"workbook_categories"`
}

// problem godoc
//
//	@Summary	API
//	@Description
//	@Tags		problem
//	@Accept		json
//	@Produce	json
//	@Param		request	body	ProblemCreateRequest	true	"問題作成API"
//	@Success	201		{object} 	nil
//	@Failure	400		{object}	app_types.ErrorResponse
//	@Failure	401		{object}	app_types.ErrorResponse
//	@Failure	500		{object}	app_types.ErrorResponse
//	@Router		/{workbook_id}/problems [post]
func (a *ProblemController) Create(c *gin.Context) {

}
