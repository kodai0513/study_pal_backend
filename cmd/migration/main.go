package main

import (
	"context"
	"log"

	"study-pal-backend/ent"
	"study-pal-backend/ent/migrate"

	_ "github.com/lib/pq"
)

func main() {
	// PostgreSQL接続
	client, err := ent.Open("postgres", "postgres://postgres:postgres@study_pal_db:5432/study_pal?search_path=public&sslmode=disable")
	if err != nil {
		log.Fatalf("failed connecting to postgresql: %v", err)
	}
	defer client.Close()

	// コンテキスト作成
	ctx := context.Background()

	// マイグレーションを実行
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	log.Println("Schema migration completed successfully.")
}
