package pages

import (
	"errors"
	"reflect"
	"strconv"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/utils/application_errors"
	"study-pal-backend/app/utils/converts"
)

func CreatePage[T any](
	baseQuery *func() ([]*T, error),
	nextQuery *func() ([]*T, error),
	prevQuery *func() ([]*T, error),
	page *app_types.Page,
	structIdIndex int,
) (
	[]*T, *app_types.Page,
	application_errors.ApplicationError,
) {
	if converts.StringToInt(page.NextPageId(), 0) > 0 && converts.StringToInt(page.PrevPageId(), 0) > 0 {
		return nil, nil, application_errors.NewClientInputValidationApplicationError(errors.New("both nextPageId and prevPageId are set. please specify only one"))
	}
	query := baseQuery
	if converts.StringToInt(page.NextPageId(), 0) > 0 {
		query = nextQuery
	}
	if converts.StringToInt(page.PrevPageId(), 0) > 0 {
		query = prevQuery
	}

	results, err := (*query)()
	if err != nil {
		return nil, nil, application_errors.NewDatabaseConnectionApplicationError(err)
	}

	nextPage := app_types.NewPage(0, "", "")
	if len(results) > page.PageSize() {
		id := reflect.ValueOf(results[len(results)-1]).Elem().Field(structIdIndex).Interface().(int)
		nextPage.SetNextPageId(strconv.Itoa(id))
		results = results[:len(results)-1]
		nextPage.SetPageSize(len(results))
	}

	nextPage.SetPageSize(len(results))

	return results, nextPage, nil
}
