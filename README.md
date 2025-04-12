# 開発方針について

### ドメイン層など整合性が必要な構造体、interfaceを通して利用する構造体、アプリケーション上でグロバールに参照する構造体に関してはに関してはコンストラクタを利用する


### 繰り返し処理はなるべく下記のライブラリで書くこと
https://github.com/samber/lo

# ent.の使用上の注意点

### クエリの発行するメソッドは○○Xという命名のメソッドを使用する(致命的なエラーをerrorの戻り値ではなくpanicによって発生させれるため)
```go
w.client.Workbook.Create().
		SetID(workbook.Id()).
		SetCreatedID(workbook.UserId()).
		SetDescription(workbook.Description()).
		SetIsPublic(workbook.IsPublic()).
		SetTitle(workbook.Title()).
		Save(w.ctx)
```

```go
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
```sh
go run -mod=mod entgo.io/ent/cmd/ent new User
```

#### entのアセットの生成
```sh
go generate ./ent
```

### マイグレーション実行
```sh
cd cmd/migration

go run main.go
```


# API仕様書の自動生成手順
下記を使用している
https://github.com/swaggo/swag

#### controllerのメソッドに以下のようにして記述
```go
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
```sh
swag fmt
```

#### swaggerファイルを自動生成
```sh
swag init
```

### APIドキュメント(swagger)
http://localhost:8080/swagger/index.html


### ローカル環境の操作

```sh
# 開発サーバーに入る
docker exec -it study_pal_backend sh

# 開発サーバーのログを表示する
docker logs study_pal_backend --tail 20 -f
```

### データベースへのアクセス
```sh
psql -h study_pal_db -p 5432 -U postgres -d study_pal
Password for user postgres: postgres
```

### controller、usecase層のテンプレートファイル作成
```sh
cd generate

# コントローラー作成
go run gen.go controller [コントローラー名]
# example
go run gen.go controller Test

# ユースケースアクション作成
go run gen.go action [アクション名]
#example
go run gen.go action Test

# ユースケースクエリ作成
go run gen.go query [クエリ名]
#example
go run gen.go query Test
```