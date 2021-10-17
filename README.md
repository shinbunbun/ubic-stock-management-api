# 設計
[hackmdで書いた設計](https://hackmd.io/Rtdli98GSKmfWD4d_DxbNA)

# 利用するソフト
 - docker-compose
  - dynamodbをローカルで立てるために使用
 - [aws cli](https://aws.amazon.com/jp/cli/)
  - dynamodbをローカルでテストするために使用
 - [golangci-lint](https://github.com/golangci/golangci-lint)
  - github actionsのバリデーションで使うので，push前に手元で通ることを確認したほうが良いです。


# テストをするために
現状はテストのためにdynamodbをローカルで立てる必要があります。以下のコマンドを実行してください。

```
docker-compose up -d # ローカルでdynamodbを立ち上げる
bash setup-database.sh # テスト用にテーブルを作成

cd stock-management-system/infrastructure/database # dynamodbを実際にテストで使うパッケージの例
go test
```
