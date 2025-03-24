package controllers

import (
	"net/http"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/controllers/shared/mappers"
	"study-pal-backend/app/infrastructures/query_services"
	timelines "study-pal-backend/app/usecases/timeline"
	"study-pal-backend/app/utils/type_converts"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type TimelineController struct {
	AppData *app_types.AppData
}

type TimelineResponse struct {
	Id           uuid.UUID `json:"id"`
	Description  string    `json:"description"`
	UserId       uuid.UUID `json:"user_id"`
	UserName     string    `json:"user_name"`
	UserNickName string    `json:"user_nick_name"`
}

type IndexResponse struct {
	PageInfo  *app_types.PageResponse `json:"page_info"`
	Timelines []*TimelineResponse     `json:"timelines"`
}

// timelines godoc
//
//	@Summary		タイムライン取得API
//	@Description	タイムラインを取得します
//	@Tags			timelines
//	@Accept			json
//	@Produce		json
//	@Param			page_size		query		int	false	"ページサイズ"
//	@Param			next_page_id	query		int	false	"次のページのid"
//	@Success		200				{object}	IndexResponse
//	@Failure		400				{object}	app_types.ErrorResponse
//	@Failure		500				{object}	app_types.ErrorResponse
//	@Router			/timelines [get]
func (t *TimelineController) Index(c *gin.Context) {
	pageSizeInput := c.Query("page_size")
	pageSize := type_converts.StringToInt(pageSizeInput, 50)

	nextPageIdInput := c.Query("next_page_id")
	timelineDtos, page, usecaseErrGroup := query_services.NewTimelineQueryServiceImpl(t.AppData.Client(), c).Fetch(&app_types.Page{
		PageSize:   pageSize,
		PrevPageId: "",
		NextPageId: nextPageIdInput,
	})

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

	timelineResponses := lo.Map(
		timelineDtos,
		func(timelineDto *timelines.TimelineDto, index int) *TimelineResponse {
			return &TimelineResponse{
				Id:           timelineDto.Id,
				Description:  timelineDto.Description,
				UserId:       timelineDto.UserId,
				UserName:     timelineDto.UserName,
				UserNickName: timelineDto.UserNickName,
			}
		},
	)
	c.SecureJSON(
		http.StatusOK,
		&IndexResponse{
			&app_types.PageResponse{
				PageSize:   page.PageSize,
				PrevPageId: page.PrevPageId,
				NextPageId: page.NextPageId,
			},
			timelineResponses,
		},
	)
}
