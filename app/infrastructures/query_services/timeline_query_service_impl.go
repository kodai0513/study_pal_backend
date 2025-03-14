package query_services

import (
	"context"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/infrastructures/query_services/shared/create_pages"
	"study-pal-backend/app/usecases/shared/usecase_errors"
	timeline_query_service "study-pal-backend/app/usecases/timelines"
	"study-pal-backend/app/utils/type_converts"
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

func (t *TimelineQueryServiceImpl) Fetch(page *app_types.Page) ([]*timeline_query_service.TimelineDto, *app_types.Page, usecase_errors.UsecaseErrorGroup) {
	limit := page.PageSize() + 1
	baseQuery := func() []*ent.Article {
		return t.client.Article.Query().WithPost().Limit(limit).AllX(t.ctx)
	}
	nextQuery := func() []*ent.Article {
		return t.client.Article.Query().WithPost().Where(article.IDGTE(type_converts.StringToInt(page.NextPageId(), 0))).Limit(limit).AllX(t.ctx)
	}

	resutls, nextPage, err := create_pages.CreatePage[ent.Article](
		&baseQuery,
		&nextQuery,
		nil,
		page,
		1,
	)

	if err != nil {
		return nil, nil, usecase_errors.NewUsecaseErrorGroupWithMessage(usecase_errors.NewUsecaseError(usecase_errors.InvalidParameter, err))
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
