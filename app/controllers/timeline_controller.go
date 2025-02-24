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

func (t *TimelineController) Index(c *gin.Context) {
	timelineList, page, err := query_services.NewTimelineQueryServiceImpl(c, t.appData.Client()).Fetch(app_types.NewPage(10, "a", "b"))

	if err != nil && err.Kind() == application_errors.ClientInputValidation {
		c.SecureJSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
	}

	if err != nil && err.Kind() == application_errors.DatabaseConnection {
		c.SecureJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
	}

	var timelineResponses []*TimelineResponse
	for _, timeline := range timelineList {
		timelineResponses = append(timelineResponses, NewTimelineResponse(timeline))
	}
	c.SecureJSON(
		http.StatusOK,
		gin.H{
			"timelines": timelineResponses,
			"page_info": app_types.NewPageResponse(page),
		},
	)
}
