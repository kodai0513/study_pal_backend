package query_services

import (
	"context"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/usecases"
	"study-pal-backend/app/utils/application_errors"
	"study-pal-backend/ent"
)

type TimelineQueryServiceImpl struct {
	ctx    context.Context
	client *ent.Client
}

func NewTimelineQueryServiceImpl(ctx context.Context, client *ent.Client) usecases.TimelineQueryService {
	return &TimelineQueryServiceImpl{
		ctx:    ctx,
		client: client,
	}
}

func (t *TimelineQueryServiceImpl) Fetch(page *app_types.Page) ([]*usecases.TimelineDto, *app_types.Page, application_errors.ApplicationError) {
	results, err := t.client.Article.Query().WithPost().All(t.ctx)

	if err != nil {
		return nil, nil, application_errors.NewDatabaseConnectionApplicationError(err)
	}

	var timelineList []*usecases.TimelineDto
	for _, result := range results {
		timelineList = append(timelineList, usecases.NewTimelineDto(
			result.ID,
			result.Description,
			result.PostID,
			result.Edges.Post.Name,
			result.Edges.Post.NickName,
		))
	}
	return timelineList, app_types.NewPage(2, "a", "b"), nil
}
