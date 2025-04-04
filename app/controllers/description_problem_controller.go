package controllers

import (
	"net/http"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/controllers/shared/mappers"
	"study-pal-backend/app/infrastructures/repositories"
	"study-pal-backend/app/usecases/description_problems"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DescriptionProblemController struct {
	AppData *app_types.AppData
}

type UpdateDescriptionProblem struct {
	CorrentStatement string `json:"correct_statement"`
	Statement        string `json:"statement"`
}

type UpdateDescriptionProblemRequest struct {
	*UpdateDescriptionProblem
}

type UpdateDescriptionProblemResponse struct {
	*UpdateDescriptionProblem
}

// descriptionproblem godoc
//
//	@Summary		API
//	@Description	問題を編集します
//	@Tags			description-problem
//	@Accept			json
//	@Produce		json
//	@Param			request					body		UpdateDescriptionProblemRequest	true	"記述問題更新リクエスト"
//	@Param			description_problem_id	path		string							true	"DescriptionProblem ID"
//	@Success		200						{object}	UpdateDescriptionProblemResponse
//	@Failure		400						{object}	app_types.ErrorResponse
//	@Failure		401						{object}	app_types.ErrorResponse
//	@Failure		500						{object}	app_types.ErrorResponse
//	@Router			/description-problems/{description_problem_id} [put]
func (a *DescriptionProblemController) Update(c *gin.Context) {
	var request UpdateDescriptionProblemRequest
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

	descriptionProblemIdParam := c.Param("description_problem_id")
	descriptionProblemId, err := uuid.Parse(descriptionProblemIdParam)
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

	action := &description_problems.UpdateAction{
		DescriptionProblemRepository: repositories.NewDescriptionProblemRepositoryImpl(a.AppData.Client(), c),
	}
	descriptionProblemDto, usecaseErrGroup := action.Execute(
		&description_problems.UpdateActionCommand{
			DescriptionProblemId: descriptionProblemId,
			CorrentStatement:     request.CorrentStatement,
			Statement:            request.Statement,
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

	c.SecureJSON(
		http.StatusOK,
		&UpdateDescriptionProblemResponse{
			&UpdateDescriptionProblem{
				CorrentStatement: descriptionProblemDto.CorrentStatement,
				Statement:        descriptionProblemDto.Statement,
			},
		},
	)
}

// descriptionproblem godoc
//
//	@Summary	API
//	@Description
//	@Tags		description-problem
//	@Accept		json
//	@Produce	json
//	@Param		description_problem_id	path		string	true	"DescriptionProblem ID"
//	@Success	204						{object}	nil
//	@Failure	400						{object}	app_types.ErrorResponse
//	@Failure	401						{object}	app_types.ErrorResponse
//	@Failure	500						{object}	app_types.ErrorResponse
//	@Router		/description-problems/{description_problem_id} [delete]
func (a *DescriptionProblemController) Delete(c *gin.Context) {
	descriptionProblemId, err := uuid.Parse(c.Param("description_problem_id"))
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

	action := &description_problems.DeleteAction{
		DescriptionProblemRepository: repositories.NewDescriptionProblemRepositoryImpl(a.AppData.Client(), c),
	}
	usecaseErrGroup := action.Execute(
		&description_problems.DeleteActionCommand{
			DescriptionProblemId: descriptionProblemId,
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

	c.SecureJSON(
		http.StatusNoContent,
		gin.H{},
	)
}
