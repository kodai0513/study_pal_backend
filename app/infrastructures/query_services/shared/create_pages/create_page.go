package create_pages

import (
	"errors"
	"reflect"
	"strconv"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/utils/type_converts"
)

func CreatePage[T any](
	baseQuery *func() []*T,
	nextQuery *func() []*T,
	prevQuery *func() []*T,
	page *app_types.Page,
	structIdIndex int,
) (
	[]*T,
	*app_types.Page,
	error,
) {
	if type_converts.StringToInt(page.NextPageId(), 0) > 0 && type_converts.StringToInt(page.PrevPageId(), 0) > 0 {
		return nil, nil, errors.New("both nextPageId and prevPageId are set. please specify only one")
	}
	query := baseQuery
	if type_converts.StringToInt(page.NextPageId(), 0) > 0 {
		query = nextQuery
	}
	if type_converts.StringToInt(page.PrevPageId(), 0) > 0 {
		query = prevQuery
	}

	results := (*query)()

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
