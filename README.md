# Todo App

Go + Nuxt.js で構築した Todo アプリケーションです。

## 技術スタック

| Layer | Technology |
|-------|------------|
| Frontend | Nuxt 3, Vue 3, Tailwind CSS |
| Backend | Go 1.21, chi router |
| Database | MySQL 8.0 |
| Infrastructure | Docker Compose |

## 起動方法

### 1. 環境変数の設定

```bash
cp .env.example .env
```

### 2. Docker Compose で起動

```bash
# マイグレーションを含めて起動
docker compose --profile tools up -d

# 通常起動
docker compose up -d
```

### 3. アクセス

- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080

## API 仕様

### Base URL
```
http://localhost:8080
```

### Endpoints

#### Todo 一覧取得
```
GET /api/todos
```

**Response:**
```json
[
  {
    "id": "uuid",
    "title": "Todo title",
    "is_completed": false,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
]
```

---

#### Todo 新規作成
```
POST /api/todos
```

**Request Body:**
```json
{
  "title": "New todo"
}
```

**Response:** `201 Created`
```json
{
  "id": "uuid",
  "title": "New todo",
  "is_completed": false,
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

---

#### Todo 完了状態更新
```
PATCH /api/todos/{id}
```

**Request Body:**
```json
{
  "is_completed": true
}
```

**Response:** `200 OK`
```json
{
  "id": "uuid",
  "title": "Todo title",
  "is_completed": true,
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

---

#### Todo 削除
```
DELETE /api/todos/{id}
```

**Response:** `204 No Content`

## 開発

### テスト実行

```bash
cd backend
go test ./...
```

### ビルド

```bash
# Backend
cd backend
go build ./...

# Frontend
cd frontend
npm run build
```
