package repositories

import (
	"context"
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/models/value_objects/workbook_categories"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/infrastructures/repositories/shared/split"
	"study-pal-backend/ent"
	"study-pal-backend/ent/workbookcategory"
	"study-pal-backend/ent/workbookcategoryclosure"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type WorkbookCategoryRepositoryImpl struct {
	tx  *ent.Tx
	ctx context.Context
}

func NewWorkbookCategoryRepositoryImpl(tx *ent.Tx, ctx context.Context) repositories.WorkbookCategoryRepository {
	return &WorkbookCategoryRepositoryImpl{
		tx:  tx,
		ctx: ctx,
	}
}

func (w *WorkbookCategoryRepositoryImpl) FindByWorkbookId(workbookId uuid.UUID) []*entities.WorkbookCategory {
	resultCategoryClosures := w.tx.WorkbookCategoryClosure.Query().
		Where(workbookcategoryclosure.WorkbookIDEQ(workbookId)).
		Order(ent.Asc(workbookcategoryclosure.FieldPosition)).
		AllX(w.ctx)
	IdToResultCategories := lo.SliceToMap(
		w.tx.WorkbookCategory.Query().Where(workbookcategory.WorkbookIDEQ(workbookId)).AllX(w.ctx),
		func(w *ent.WorkbookCategory) (uuid.UUID, *ent.WorkbookCategory) {
			return w.ID, w
		},
	)

	categoryEntities := lo.FilterMap(resultCategoryClosures, func(w *ent.WorkbookCategoryClosure, _ int) (*entities.WorkbookCategory, bool) {
		if !w.IsRoot {
			return nil, false
		}

		name, _ := workbook_categories.NewName(IdToResultCategories[w.ChildID].Name)
		return entities.CreateWorkbookCategory(
			IdToResultCategories[w.ChildID].ID,
			name,
			IdToResultCategories[w.ChildID].WorkbookID,
		), true
	})

	lo.ForEach(categoryEntities, func(w *entities.WorkbookCategory, _ int) {
		createResult(resultCategoryClosures, w, IdToResultCategories, w.Id())
	})

	return categoryEntities
}

func (w *WorkbookCategoryRepositoryImpl) UpsertAndDeleteBulk(workbookCategories []*entities.WorkbookCategory, workbookId uuid.UUID) []*entities.WorkbookCategory {
	var createData func(parentIds []uuid.UUID, position *int, category *entities.WorkbookCategory) ([]*ent.WorkbookCategory, []*ent.WorkbookCategoryClosure)
	createData = func(parentIds []uuid.UUID, position *int, category *entities.WorkbookCategory) ([]*ent.WorkbookCategory, []*ent.WorkbookCategoryClosure) {
		var entCategories []*ent.WorkbookCategory
		var entCategoryClosures []*ent.WorkbookCategoryClosure

		entCategories = append(entCategories, &ent.WorkbookCategory{
			ID:         category.Id(),
			Name:       category.Name(),
			WorkbookID: category.WorkbookId(),
		})

		if len(parentIds) > 0 {
			entCategoryClosures = append(entCategoryClosures, lo.Map(parentIds, func(parentId uuid.UUID, i int) *ent.WorkbookCategoryClosure {
				return &ent.WorkbookCategoryClosure{
					ID:         uuid.New(),
					WorkbookID: category.WorkbookId(),
					ChildID:    category.Id(),
					IsRoot:     false,
					ParentID:   parentId,
					Position:   *position,
					Level:      len(parentIds) - i,
				}
			})...)

			entCategoryClosures = append(entCategoryClosures, &ent.WorkbookCategoryClosure{
				ID:         uuid.New(),
				WorkbookID: category.WorkbookId(),
				ChildID:    category.Id(),
				IsRoot:     false,
				ParentID:   category.Id(),
				Position:   *position,
				Level:      0,
			})
		} else {
			entCategoryClosures = append(entCategoryClosures, &ent.WorkbookCategoryClosure{
				ID:         uuid.New(),
				WorkbookID: category.WorkbookId(),
				ChildID:    category.Id(),
				IsRoot:     true,
				ParentID:   category.Id(),
				Position:   *position,
				Level:      0,
			})
		}

		for _, children := range category.Children() {
			*position = *position + 1
			childCategories, childCategoryClosures := createData(append(parentIds, category.Id()), position, children)
			entCategories = append(entCategories, childCategories...)
			entCategoryClosures = append(entCategoryClosures, childCategoryClosures...)
		}

		return entCategories, entCategoryClosures
	}

	position := 0
	var newCategories []*ent.WorkbookCategory
	var newCategoryClosures []*ent.WorkbookCategoryClosure
	lo.ForEach(workbookCategories, func(category *entities.WorkbookCategory, _ int) {
		categories, categoryClosures := createData([]uuid.UUID{}, &position, category)
		newCategories = append(newCategories, categories...)
		newCategoryClosures = append(newCategoryClosures, categoryClosures...)
	})

	split := split.UpsertAndDeleteSplit(
		lo.Map(newCategories, func(category *ent.WorkbookCategory, _ int) uuid.UUID {
			return category.ID
		}),
		w.tx.WorkbookCategory.Query().Where(workbookcategory.WorkbookIDEQ(workbookId)).IDsX(w.ctx),
	)

	IdToNewCategories := lo.SliceToMap(newCategories, func(w *ent.WorkbookCategory) (uuid.UUID, *ent.WorkbookCategory) {
		return w.ID, w
	})
	w.tx.WorkbookCategory.MapCreateBulk(
		lo.FilterMap(split.CreateIds, func(createId uuid.UUID, _ int) (*ent.WorkbookCategory, bool) {
			createCategory, ok := IdToNewCategories[createId]
			return createCategory, ok
		}),
		func(c *ent.WorkbookCategoryCreate, i int) {
			c.
				SetID(newCategories[i].ID).
				SetName(newCategories[i].Name).
				SetWorkbookID(newCategories[i].WorkbookID)
		}).SaveX(w.ctx)
	w.tx.WorkbookCategoryClosure.Delete().Where(workbookcategoryclosure.WorkbookIDEQ(workbookId)).ExecX(w.ctx)
	w.tx.WorkbookCategory.Delete().Where(workbookcategory.IDIn(split.DeleteIds...)).ExecX(w.ctx)
	lo.ForEach(split.UpdateIds, func(updateId uuid.UUID, _ int) {
		updateCategory := IdToNewCategories[updateId]
		w.tx.WorkbookCategory.UpdateOneID(updateId).
			SetName(updateCategory.Name).
			SaveX(w.ctx)
	})

	w.tx.WorkbookCategoryClosure.MapCreateBulk(newCategoryClosures, func(c *ent.WorkbookCategoryClosureCreate, i int) {
		c.
			SetID(newCategoryClosures[i].ID).
			SetWorkbookID(newCategoryClosures[i].WorkbookID).
			SetChildID(newCategoryClosures[i].ChildID).
			SetIsRoot(newCategoryClosures[i].IsRoot).
			SetParentID(newCategoryClosures[i].ParentID).
			SetPosition(newCategoryClosures[i].Position).
			SetLevel(newCategoryClosures[i].Level)
	}).SaveX(w.ctx)

	resultCategoryClosures := w.tx.WorkbookCategoryClosure.Query().
		Where(workbookcategoryclosure.WorkbookIDEQ(workbookId)).
		Order(ent.Asc(workbookcategoryclosure.FieldPosition)).
		AllX(w.ctx)
	IdToResultCategories := lo.SliceToMap(
		w.tx.WorkbookCategory.Query().Where(workbookcategory.WorkbookIDEQ(workbookId)).AllX(w.ctx),
		func(w *ent.WorkbookCategory) (uuid.UUID, *ent.WorkbookCategory) {
			return w.ID, w
		},
	)

	resultCategoryEntities := lo.FilterMap(resultCategoryClosures, func(w *ent.WorkbookCategoryClosure, _ int) (*entities.WorkbookCategory, bool) {
		if !w.IsRoot {
			return nil, false
		}

		name, _ := workbook_categories.NewName(IdToResultCategories[w.ChildID].Name)
		return entities.CreateWorkbookCategory(
			IdToResultCategories[w.ChildID].ID,
			name,
			IdToResultCategories[w.ChildID].WorkbookID,
		), true
	})

	lo.ForEach(resultCategoryEntities, func(w *entities.WorkbookCategory, _ int) {
		createResult(resultCategoryClosures, w, IdToResultCategories, w.Id())
	})

	return resultCategoryEntities
}

func createResult(
	categoryClosures []*ent.WorkbookCategoryClosure,
	categoryEntity *entities.WorkbookCategory,
	IdToCategories map[uuid.UUID]*ent.WorkbookCategory,
	parentId uuid.UUID,
) {
	children := lo.FilterMap(categoryClosures, func(closure *ent.WorkbookCategoryClosure, _ int) (*entities.WorkbookCategory, bool) {
		// 子のデータだけを取得したい
		if closure.ParentID != parentId || closure.Level != 1 {
			return nil, false
		}

		name, _ := workbook_categories.NewName(IdToCategories[closure.ChildID].Name)
		return entities.CreateWorkbookCategory(
			IdToCategories[closure.ChildID].ID,
			name,
			IdToCategories[closure.ChildID].WorkbookID,
		), true
	})

	lo.ForEach(children, func(child *entities.WorkbookCategory, _ int) {
		categoryEntity.AddChild(child)
		createResult(categoryClosures, child, IdToCategories, child.Id())
	})
}
