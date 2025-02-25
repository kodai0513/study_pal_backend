package controllers

import (
	"net/http"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/infrastructures/query_services"
	"study-pal-backend/app/usecases"
	"study-pal-backend/app/utils/application_errors"

	"github.com/gin-gonic/gin"
)

type TimelineController struct {
	appData *app_types.AppData
}

func NewTimelineController(appData *app_types.AppData) *TimelineController {
	return &TimelineController{
		appData: appData,
	}
}

type TimelineResponse struct {
	Id           int    `json:"id"`
	Description  string `json:"description"`
	PostId       int    `json:"post_id"`
	PostName     string `json:"post_name"`
	PostNickName string `json:"post_nick_name"`
}

func NewTimelineResponse(timeline *usecases.TimelineDto) *TimelineResponse {
	return &TimelineResponse{
		Id:           timeline.Id(),
		Description:  timeline.Description(),
		PostId:       timeline.PostId(),
		PostName:     timeline.PostName(),
		PostNickName: timeline.PostNickName(),
	}
}

type IndexResponse struct {
	PageInfo  *app_types.PageResponse `json:"page_info"`
	Timelines []*TimelineResponse     `json:"timelines"`
}

func NewIndexResponse(pageResponse *app_types.PageResponse, timelineResponses []*TimelineResponse) *IndexResponse {
	return &IndexResponse{
		PageInfo:  pageResponse,
		Timelines: timelineResponses,
	}
}

// timelines godoc
//
//	@Summary		タイムライン取得API
//	@Description	タイムラインを取得します
//	@Tags			timelines
//	@Accept			json
//	@Produce		json
//	@Param			page_size		query		int		true "ページサイズ"
//	@Param			prev_page_token	query		string	false "次のページのトークン"
//	@Param			next_page_token	query		string	false "前のページのトークン"
//	@Success		200				{object}	IndexResponse
//	@Failure		400				{object}	app_types.ErrorResponse
//	@Failure		500				{object}	app_types.ErrorResponse
//	@Router			/timelines [get]
func (t *TimelineController) Index(c *gin.Context) {
	_ = c.Query("page_size")
	timelineList, page, err := query_services.NewTimelineQueryServiceImpl(c, t.appData.Client()).Fetch(app_types.NewPage(10, "a", "b"))

	if err != nil && err.Kind() == application_errors.ClientInputValidation {
		var errors []string
		errors = append(errors, err.Error())
		c.SecureJSON(
			http.StatusBadRequest,
			app_types.NewErrorResponse(errors),
		)
	}

	if err != nil && err.Kind() == application_errors.DatabaseConnection {
		var errors []string
		errors = append(errors, err.Error())
		c.SecureJSON(
			http.StatusInternalServerError,
			app_types.NewErrorResponse(errors),
		)
	}

	var timelineResponses []*TimelineResponse
	for _, timeline := range timelineList {
		timelineResponses = append(timelineResponses, NewTimelineResponse(timeline))
	}
	c.SecureJSON(
		http.StatusOK,
		NewIndexResponse(app_types.NewPageResponse(page), timelineResponses),
	)
}
