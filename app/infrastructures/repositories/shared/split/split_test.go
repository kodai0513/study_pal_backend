package split

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUpsertAndDeleteSplit_正常にスプリットが作成されるか(t *testing.T) {
	// テスト用のUUIDを生成
	id1 := uuid.New()
	id2 := uuid.New()
	id3 := uuid.New()
	id4 := uuid.New()

	// 新規IDと登録済みIDのセットを作成
	newIds := []uuid.UUID{id1, id2, id3}
	registeredIds := []uuid.UUID{id2, id3, id4}

	// UpsertAndDeleteSplit関数を実行
	split := UpsertAndDeleteSplit(newIds, registeredIds)

	// CreateIdsには新規に追加されたIDが含まれることを確認
	assert.ElementsMatch(t, split.CreateIds, []uuid.UUID{id1})

	// DeleteIdsには削除されたIDが含まれることを確認
	assert.ElementsMatch(t, split.DeleteIds, []uuid.UUID{id4})

	// UpdateIdsには更新されたIDが含まれることを確認
	assert.ElementsMatch(t, split.UpdateIds, []uuid.UUID{id2, id3})
}

func TestUpsertAndDeleteSplit_新規IDだけの場合(t *testing.T) {
	// 新規IDだけのセット
	newIds := []uuid.UUID{uuid.New(), uuid.New()}
	registeredIds := []uuid.UUID{}

	// UpsertAndDeleteSplit関数を実行
	split := UpsertAndDeleteSplit(newIds, registeredIds)

	// CreateIdsには新規IDがすべて含まれることを確認
	assert.ElementsMatch(t, split.CreateIds, newIds)

	// DeleteIdsとUpdateIdsは空であることを確認
	assert.Empty(t, split.DeleteIds)
	assert.Empty(t, split.UpdateIds)
}

func TestUpsertAndDeleteSplit_登録IDだけの場合(t *testing.T) {
	// 登録IDだけのセット
	newIds := []uuid.UUID{}
	registeredIds := []uuid.UUID{uuid.New(), uuid.New()}

	// UpsertAndDeleteSplit関数を実行
	split := UpsertAndDeleteSplit(newIds, registeredIds)

	// CreateIdsとUpdateIdsは空であることを確認
	assert.Empty(t, split.CreateIds)
	assert.Empty(t, split.UpdateIds)

	// DeleteIdsには登録IDがすべて含まれることを確認
	assert.ElementsMatch(t, split.DeleteIds, registeredIds)
}

func TestUpsertAndDeleteSplit_すべてのIDが一致しない場合(t *testing.T) {
	// 新規IDと登録IDがまったく重ならないケース
	newIds := []uuid.UUID{uuid.New(), uuid.New()}
	registeredIds := []uuid.UUID{uuid.New(), uuid.New()}

	// UpsertAndDeleteSplit関数を実行
	split := UpsertAndDeleteSplit(newIds, registeredIds)

	// CreateIdsには新規IDがすべて含まれることを確認
	assert.ElementsMatch(t, split.CreateIds, newIds)

	// DeleteIdsには登録IDがすべて含まれることを確認
	assert.ElementsMatch(t, split.DeleteIds, registeredIds)

	// UpdateIdsは空であることを確認
	assert.Empty(t, split.UpdateIds)
}
