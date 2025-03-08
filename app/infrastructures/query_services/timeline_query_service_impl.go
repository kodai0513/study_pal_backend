package query_services

import (
	"context"
	"study-pal-backend/app/app_types"
	timeline_query_service "study-pal-backend/app/usecases/timelines"
	"study-pal-backend/app/utils/application_errors"
	"study-pal-backend/app/utils/converts"
	"study-pal-backend/app/utils/pages"
	"study-pal-backend/ent"
	"study-pal-backend/ent/article"
)

type TimelineQueryServiceImpl struct {
	ctx    context.Context
	client *ent.Client
}

func NewTimelineQueryServiceImpl(ctx context.Context, client *ent.Client) timeline_query_service.TimelineQueryService {
	return &TimelineQueryServiceImpl{
		ctx:    ctx,
		client: client,
	}
}

func (t *TimelineQueryServiceImpl) Fetch(page *app_types.Page) ([]*timeline_query_service.TimelineDto, *app_types.Page, application_errors.ApplicationError) {
	limit := page.PageSize() + 1
	baseQuery := func() ([]*ent.Article, error) {
		return t.client.Article.Query().WithPost().Limit(limit).All(t.ctx)
	}
	nextQuery := func() ([]*ent.Article, error) {
		return t.client.Article.Query().WithPost().Where(article.IDGTE(converts.StringToInt(page.NextPageId(), 0))).Limit(limit).All(t.ctx)
	}

	resutls, nextPage, err := pages.CreatePage[ent.Article](
		&baseQuery,
		&nextQuery,
		nil,
		page,
		1,
	)

	if err != nil {
		return nil, nil, err
	}

	var timelineDtos []*timeline_query_service.TimelineDto
	for _, result := range resutls {
		timelineDtos = append(timelineDtos, timeline_query_service.NewTimelineDto(
			result.ID,
			result.Description,
			result.PostID,
			result.Edges.Post.Name,
			result.Edges.Post.NickName,
		))
	}

	return timelineDtos, nextPage, nil
}
