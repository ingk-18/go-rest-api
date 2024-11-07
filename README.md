# 環境構築
docker-compose upでコンテナ起動後、以下コマンドでテーブルを作成する
1. $env:GO_ENV = "dev"  
2. go run .\migrate\migrate.go