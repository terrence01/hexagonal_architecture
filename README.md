# 六角架構 (Hexagonal Architecture) 範例

本專案展示了一個使用六角架構（端口和適配器模式）實現的用戶管理系統，包含創建用戶和刪除用戶的用例。使用 Gin 框架作為 HTTP 服務器。

## 專案結構

```
internal/
└── core/                 # 核心業務邏輯與適配器（一個限界上下文）
    ├── domain/           # 領域模型
    │   ├── user.go
    │   └── user_repository.go  # 用戶存儲庫接口（領域層）
    ├── ports/            # 端口（接口）
    │   ├── in/           # 輸入端口（用例）
    │   │   ├── command/  # 命令對象
    │   │   │   └── create_user_command.go
    │   │   └── user_usecase.go
    │   └── out/          # 輸出端口（外部服務）
    │       └── email_sender.go
    ├── application/      # 應用服務（用例實現）
    │   └── user/         # 用戶相關服務
    │       ├── common.go
    │       ├── create_user_service.go
    │       ├── delete_user_service.go
    │       ├── get_user_service.go
    │       └── user_service_factory.go
    └── adapters/         # 適配器（與外部世界的連接）
        ├── in/           # 輸入適配器
        │   ├── controller/   # 控制器
        │   │   └── user/     # 用戶相關控制器
        │   │       ├── create_user_controller.go
        │   │       ├── delete_user_controller.go
        │   │       ├── get_user_controller.go
        │   │       └── user_controller_factory.go
        │   └── http/         # HTTP 適配器
        │       ├── dto/      # 數據傳輸對象
        │       │   └── user.go
        │       └── router/   # 路由配置
        │           └── gin_router.go
        └── out/              # 輸出適配器
            ├── email/        # 電子郵件實現
            │   └── email_sender.go
            └── repository/   # 數據庫實現
                └── memory/   # 內存存儲實現
                    └── user_repository.go
```

## 使用說明

本專案展示了六角架構模式，其中：
- Domain（領域）包含業務實體和存儲庫接口
  - User：用戶實體
  - UserRepository：用戶存儲庫接口（領域驅動設計中的存儲庫模式）
- Ports（端口）定義了輸入（用例）和輸出（外部服務）的接口
  - Command：用例的輸入參數對象
  - 輸出端口：
    - EmailSender：發送電子郵件接口
- Application（應用服務）實現了用例
  - 每個用例都有獨立的服務實現文件
  - 所有服務都實現了標準的用例接口
- Adapters（適配器）實現了與外部系統連接的接口
  - 輸入適配器：
    - 控制器：每個用例都有獨立的控制器文件
    - DTO：數據傳輸對象，用於請求和響應
    - 路由：連接框架特定路由與控制器
  - 輸出適配器：
    - 數據庫實現（按存儲類型分類，如內存存儲）
    - 電子郵件發送實現

## 限界上下文與適配器

在這個架構中，我們將適配器放在核心（core）文件夾下，這表示：
- 每個限界上下文（bounded context）擁有自己的適配器
- 適配器是特定於某個業務領域的，而不是共享的
- 這種結構更適合微服務架構，每個微服務可以有自己的核心和適配器

這種設計使得每個限界上下文可以：
- 獨立開發、測試和部署
- 選擇最適合其需求的技術實現
- 避免與其他限界上下文的不必要耦合

## 檔案命名規則

為了提高可讀性和可維護性，本專案採用了以下命名規則：
- 按職責命名文件，而不是按實體命名
- 例如：`create_user_service.go` 而不是 `user.go`
- 這樣可以更容易找到與特定用例相關的代碼

## 通訊協議擴展性

該架構設計允許輕鬆添加新的通訊協議：
- 控制器與特定通訊協議無關
- 路由器負責將特定協議（如HTTP/Gin）連接到控制器
- 要添加新協議（如GraphQL或gRPC），只需創建新的路由器實現

## API 端點

使用 Gin 框架實現的 RESTful API：

- `POST /users` - 創建新用戶
- `GET /users/:id` - 獲取用戶信息
- `DELETE /users/:id` - 刪除用戶

## 運行示例

```bash
go run cmd/core/main.go
```

訪問 http://localhost:8081 查看 API
