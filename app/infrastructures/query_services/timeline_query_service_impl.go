package query_services

import (
	"context"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/infrastructures/query_services/shared/create_pages"
	"study-pal-backend/app/usecases/shared/usecase_error"
	timeline_query_service "study-pal-backend/app/usecases/timeline"
	"study-pal-backend/app/utils/type_converts"
	"study-pal-backend/ent"
	"study-pal-backend/ent/article"

	"github.com/samber/lo"
)

type TimelineQueryServiceImpl struct {
	client *ent.Client
	ctx    context.Context
}

func NewTimelineQueryServiceImpl(client *ent.Client, ctx context.Context) timeline_query_service.TimelineQueryService {
	return &TimelineQueryServiceImpl{
		client: client,
		ctx:    ctx,
	}
}

func (t *TimelineQueryServiceImpl) Fetch(page *app_types.Page) ([]*timeline_query_service.TimelineDto, *app_types.Page, usecase_error.UsecaseErrorGroup) {
	limit := page.PageSize + 1
	baseQuery := func() []*ent.Article {
		return t.client.Article.Query().WithPost().Limit(limit).AllX(t.ctx)
	}
	nextQuery := func() []*ent.Article {
		return t.client.Article.Query().WithPost().Where(article.PageIDGTE(type_converts.StringToInt(page.NextPageId, 0))).Limit(limit).AllX(t.ctx)
	}

	results, nextPage, err := create_pages.CreatePage(
		&baseQuery,
		&nextQuery,
		nil,
		page,
		4,
	)

	if err != nil {
		return nil, nil, usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
	}

	timelineDtos := lo.Map(
		results,
		func(article *ent.Article, index int) *timeline_query_service.TimelineDto {
			return &timeline_query_service.TimelineDto{
				Id:           article.ID,
				Description:  article.Description,
				UserId:       article.UserID,
				UserNickName: article.Edges.Post.Name,
				UserName:     article.Edges.Post.NickName,
			}
		},
	)

	return timelineDtos, nextPage, nil
}
