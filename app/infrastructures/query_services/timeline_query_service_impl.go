package query_services

import (
	"context"
	"strconv"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/usecases"
	"study-pal-backend/app/utils/application_errors"
	"study-pal-backend/app/utils/converts"
	"study-pal-backend/ent"
	"study-pal-backend/ent/article"
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
	limit := page.PageSize() + 1

	query := t.client.Article.Query().WithPost().Limit(limit)

	if converts.StringToInt(page.NextPageId(), 0) > 0 {
		query = t.client.Article.Query().WithPost().Where(article.IDGTE(converts.StringToInt(page.NextPageId(), 0))).Limit(limit)
	}

	timelines, err := query.All(t.ctx)

	if err != nil {
		return nil, nil, application_errors.NewDatabaseConnectionApplicationError(err)
	}

	nextPage := app_types.NewPage(0, "", "")
	if len(timelines) > page.PageSize() {
		nextPage.SetNextPageId(strconv.Itoa(timelines[len(timelines)-1].ID))
		timelines = timelines[:len(timelines)-1]
		nextPage.SetPageSize(len(timelines))
	} else {
		nextPage.SetPageSize(len(timelines))
	}

	var timelineList []*usecases.TimelineDto
	for _, result := range timelines {
		timelineList = append(timelineList, usecases.NewTimelineDto(
			result.ID,
			result.Description,
			result.PostID,
			result.Edges.Post.Name,
			result.Edges.Post.NickName,
		))
	}
	return timelineList, nextPage, nil
}
