# attend manager api
API server

フラットな感じで実装していく

./
|---- config.go         設定関連
|---- handler.go        その名の通りハンドラー
|---- middleware.go     loggerとかミドルウェア関連
|---- model.go          データモデル
|---- service.go        ビジネスロジック
|---- db.go             永続化層。データアクセスを行う
|---- main.go           エントリポイント