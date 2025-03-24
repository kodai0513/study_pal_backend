# 開発方針について

### ドメイン層など整合性が必要な構造体、interfaceを通して利用する構造体、アプリケーション上でグロバールに参照する構造体に関してはに関してはコンストラクタを利用する
```
type Article struct {
	id          uuid.UUID
	description articles.Description
	userId      uuid.UUID
}

func NewArticle(id uuid.UUID, description articles.Description, userId uuid.UUID) *Article {
	return &Article{
		id:          id,
		description: description,
		userId:      userId,
	}
}

func (a *Article) Id() uuid.UUID {
	return a.id
}

func (a *Article) Description() string {
	return a.description.Value()
}

func (a *Article) UserId() uuid.UUID {
	return a.userId
}

```

### 繰り返し処理はなるべき下記のライブラリで書くこと
https://github.com/samber/lo

# ent.の使用上の注意点

### クエリの発行するメソッドは○○Xという命名のメソッドを使用する(致命的なエラーをerrorの戻り値ではなくpanicによって発生させれるため)
```
w.client.Workbook.Create().
		SetID(workbook.Id()).
		SetCreatedID(workbook.UserId()).
		SetDescription(workbook.Description()).
		SetIsPublic(workbook.IsPublic()).
		SetTitle(workbook.Title()).
		Save(w.ctx)
```

```
w.client.Workbook.Create().
		SetID(workbook.Id()).
		SetCreatedID(workbook.UserId()).
		SetDescription(workbook.Description()).
		SetIsPublic(workbook.IsPublic()).
		SetTitle(workbook.Title()).
		SaveX(w.ctx)
```

# マイグレーション手順

### コマンドを実行しスキーマを定義
```
go run -mod=mod entgo.io/ent/cmd/ent new User
```

#### entのアセットの生成
```
go generate ./ent
```

#### atlasgoが入っていない場合はインストールしておく
```
curl -sSf https://atlasgo.sh | sh
```

#### マイグレーションファイルの生成
```
atlas migrate diff migration_name \
    --dir "file://db/migrations" \
    --to "ent://ent/schema" \
    --dev-url "docker://postgres/15/study_pal?search_path=public"
```

#### マイグレーションの適用
```
atlas migrate apply \
    --dir "file://db/migrations" \
    --url "postgres://postgres:postgres@study_pal_db:5432/study_pal?search_path=public&sslmode=disable"
```


# API仕様書の自動生成手順
下記を使用している
https://github.com/swaggo/swag

#### controllerのメソッドに以下のようにして記述
```
// timelines godoc
//
//	@Summary		タイムライン取得API
//	@Description	タイムラインを取得します
//	@Tags			timelines
//	@Accept			json
//	@Produce		json
//	@Param			page_size		query		int		true "ページサイズ"
//	@Param			prev_page_token	query		string	false "次のページのトークン"
//	@Param			next_page_token	query		string	false "前のページのトークン"
//	@Success		200				{object}	IndexResponse
//	@Failure		400				{object}	app_types.ErrorResponse
//	@Failure		500				{object}	app_types.ErrorResponse
//	@Router			/timelines [get]
func (t *TimelineController) Index(c *gin.Context)
```

#### ファイルを自動整形
'''
swag fmt
'''

#### swaggerファイルを自動生成
```
swag init
```

### APIドキュメント(swagger)
http://localhost:8080/swagger/index.html