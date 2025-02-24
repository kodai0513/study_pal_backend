# マイグレーション手順

### entのアセットの生成
go generate ./ent

### atlasgoが入っていない場合はインストールしておく
curl -sSf https://atlasgo.sh | sh

### マイグレーションファイルの生成
atlas migrate diff migration_name \
    --dir "file://db/migrations" \
    --to "ent://ent/schema" \
    --dev-url "docker://postgres/15/study_pal?search_path=public"

### マイグレーションの適用
atlas migrate apply \
    --dir "file://db/migrations" \
    --url "postgres://postgres:postgres@study_pal_db:5432/study_pal?search_path=public&sslmode=disable"