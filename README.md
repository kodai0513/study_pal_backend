# マイグレーション手順

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