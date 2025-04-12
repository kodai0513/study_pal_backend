package split

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Split struct {
	CreateIds []uuid.UUID
	DeleteIds []uuid.UUID
	UpdateIds []uuid.UUID
}

func UpsertAndDeleteSplit(newIds []uuid.UUID, registeredIds []uuid.UUID) *Split {
	newIdMaps := lo.SliceToMap(newIds, func(id uuid.UUID) (uuid.UUID, uuid.UUID) {
		return id, id
	})
	registeredIdMaps := lo.SliceToMap(registeredIds, func(id uuid.UUID) (uuid.UUID, uuid.UUID) {
		return id, id
	})
	split := &Split{
		CreateIds: make([]uuid.UUID, 0),
		DeleteIds: make([]uuid.UUID, 0),
		UpdateIds: make([]uuid.UUID, 0),
	}

	split.CreateIds = lo.Filter(newIds, func(newId uuid.UUID, _ int) bool {
		_, ok := registeredIdMaps[newId]
		return !ok
	})
	split.DeleteIds = lo.Filter(registeredIds, func(reigsteredId uuid.UUID, _ int) bool {
		_, ok := newIdMaps[reigsteredId]
		return !ok
	})

	split.UpdateIds = lo.Filter(registeredIds, func(registeredId uuid.UUID, _ int) bool {
		_, ok := newIdMaps[registeredId]
		return ok
	})

	return split
}
