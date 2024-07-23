# Invoice Payment Service

## 概要
`invoice-payment-service`は、企業向けに請求書のデータを管理し、未来の支払期日に自動で支払いを行うことができるREST APIです。ユーザーは請求書データを登録し、支払期日に自動的に支払い処理を行うことができます。

## 特徴
- 新しい請求書データの作成
- 指定期間内に支払いが発生する請求書データの一覧取得

## ディレクトリ構造
```
/invoice-payment-service
├── /cmd
│ ├── main.go
├── /server
│ ├── /api
│ │ └── router.go
│ ├── /controllers
│ │ ├── invoice_controller.go
│ │ └── /adapter
│ │ └── response_invoice.go
│ ├── /models
│ │ ├── client.go
│ │ ├── client_bank_account.go
│ │ ├── company.go
│ │ ├── invoice.go
│ │ └── user.go
│ ├── /repository
│ │ └── invoice_repository.go
│ └── /usecase
│ └── invoice_usecase.go
```


## エンドポイント

### POST `/api/invoices`
新しい請求書データを作成します。請求金額は自動的に計算されます。

リクエスト例:
```json
{
    "issue_date": "2024-07-23",
    "payment_amount": 10000,
    "payment_due_date": "2024-08-23",
    "status": "Pending",
    "company_id": 1,
    "client_id": 1
}
```

レスポンス例:
```
{
    "id": 1,
    "issue_date": "2024-07-23",
    "payment_amount": 10000,
    "fee": 400,
    "fee_rate": 0.04,
    "sales_tax": 40,
    "sales_tax_rate": 0.1,
    "invoice_amount": 10440,
    "payment_due_date": "2024-08-23",
    "status": "Pending",
    "company_id": 1,
    "client_id": 1
}
```

### GET /api/invoices
指定期間内に支払いが発生する請求書データの一覧を取得します。

リクエスト例:
```
GET /api/invoices?start_date=2024-07-01&end_date=2024-07-31
```
レスポンス例:
```json
[
    {
        "id": 1,
        "issue_date": "2024-07-23",
        "payment_amount": 10000,
        "fee": 400,
        "fee_rate": 0.04,
        "sales_tax": 40,
        "sales_tax_rate": 0.1,
        "invoice_amount": 10440,
        "payment_due_date": "2024-08-23",
        "status": "Pending",
        "company_id": 1,
        "client_id": 1
    },
    {
        "id": 2,
        "issue_date": "2024-07-25",
        "payment_amount": 20000,
        "fee": 800,
        "fee_rate": 0.04,
        "sales_tax": 80,
        "sales_tax_rate": 0.1,
        "invoice_amount": 20880,
        "payment_due_date": "2024-08-31",
        "status": "Pending",
        "company_id": 1,
        "client_id": 1
    }
]

```

### インストールとセットアップ
1. リポジトリをクローンします
```
git clone https://github.com/shota612/invoice-payment-service.git
```
2. 必要な依存関係をインストールします
```
cd invoice-payment-service
go mod download
```
3. サーバーを起動します
```
go run cmd/main.go
```

### テスト
テストを実行するには、各レイヤごとに以下のコマンドを実行します。
```
go test ./server/controllers/tests
go test ./server/models/tests
go test ./server/repository/tests
go test ./server/usecase/tests
```