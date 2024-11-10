# 環境構築

docker-compose up でコンテナ起動後、以下コマンドでテーブルを作成する

1. $env:GO_ENV = "dev"
    - 開発環境環境変数の定義
        - 設定されているか確認
            - echo $env:GO_ENV
2. go run .\migrate\migrate.go
    - マイグレーションファイルの実行
3. go run main.go
    - アーキテクチャベースのコードが書けたらプログラムの実行
