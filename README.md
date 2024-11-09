# 環境構築
docker-compose upでコンテナ起動後、以下コマンドでテーブルを作成する
1. $env:GO_ENV = "dev"
   - 開発環境環境変数の定義
2. go run .\migrate\migrate.go
   - マイグレーションファイルの実行
3. go run main.go
   - アーキテクチャベースのコードが書けたらプログラムの実行