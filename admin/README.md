# 螃蟹商城管理后台 API

这是螃蟹商城管理后台的API服务，使用Golang和Gin框架开发，数据库为MySQL。

## 功能特性

- 用户认证（登录、注册）
- JWT令牌验证
- 基于角色的访问控制
- 数据库迁移工具

## 目录结构

```
admin/
├── cmd/                    # 命令行工具
│   └── migrate/            # 数据库迁移工具
├── config/                 # 配置相关
├── internal/               # 内部包
│   ├── auth/               # 认证相关
│   ├── controller/         # 控制器
│   ├── database/           # 数据库相关
│   ├── middleware/         # 中间件
│   ├── models/             # 数据模型
│   ├── routes/             # 路由配置
│   └── service/            # 业务逻辑
├── migrations/             # 数据库迁移文件
├── .env                    # 环境变量配置
├── go.mod                  # Go模块定义
└── main.go                 # 程序入口
```

## 开始使用

### 前置条件

- Go 1.19+
- MySQL 5.7+

### 安装

1. 克隆项目

```bash
git clone https://github.com/your-repo/crabshop-admin.git
cd crabshop-admin/admin
```

2. 安装依赖

```bash
go mod tidy
```

3. 配置环境变量

编辑 `.env` 文件，填写正确的数据库连接信息。

4. 创建数据库

```sql
CREATE DATABASE IF NOT EXISTS crabshop DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

5. 运行数据库迁移

```bash
go run cmd/migrate/main.go -direction=up
```

6. 启动服务器

```bash
go run main.go
```

## API 接口

### 认证相关

#### 登录

- **URL**: `/api/login`
- **方法**: `POST`
- **请求体**:
  ```json
  {
    "username": "admin",
    "password": "password"
  }
  ```
- **成功响应** (200 OK):
  ```json
  {
    "message": "登录成功",
    "data": {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "user": {
        "id": 1,
        "username": "admin",
        "name": "管理员",
        "email": "admin@example.com",
        "role": "admin"
      }
    }
  }
  ```

#### 注册

- **URL**: `/api/register`
- **方法**: `POST`
- **请求体**:
  ```json
  {
    "username": "newuser",
    "password": "password",
    "name": "新用户",
    "email": "user@example.com"
  }
  ```
- **成功响应** (201 Created):
  ```json
  {
    "message": "注册成功",
    "data": {
      "user": {
        "id": 2,
        "username": "newuser",
        "name": "新用户",
        "email": "user@example.com",
        "role": "user"
      }
    }
  }
  ```

#### 获取用户信息

- **URL**: `/api/user/profile`
- **方法**: `GET`
- **请求头**: `Authorization: Bearer {token}`
- **成功响应** (200 OK):
  ```json
  {
    "data": {
      "user": {
        "id": 1,
        "username": "admin",
        "name": "管理员",
        "email": "admin@example.com",
        "role": "admin"
      }
    }
  }
  ```

## 数据库迁移

使用golang-migrate库实现数据库迁移功能。

### 创建新的迁移

迁移文件命名格式为 `<version>_<name>.<direction>.sql`，例如：

- `000001_create_users_table.up.sql`
- `000001_create_users_table.down.sql`

### 应用迁移

向上迁移（创建）:

```bash
go run cmd/migrate/main.go -direction=up
```

向下迁移（回滚）:

```bash
go run cmd/migrate/main.go -direction=down
```

## 配置说明

在 `.env` 文件中可以配置以下参数：

- `PORT`: 服务器端口
- `ENV`: 环境（development/production）
- `GIN_MODE`: Gin模式（debug/release）
- `DB_HOST`: 数据库主机
- `DB_PORT`: 数据库端口
- `DB_USER`: 数据库用户名
- `DB_PASSWORD`: 数据库密码
- `DB_NAME`: 数据库名称
- `JWT_SECRET`: JWT密钥
- `JWT_EXPIRATION`: JWT过期时间（例如：24h）

## 技术栈

- [Gin](https://github.com/gin-gonic/gin): Web框架
- [GORM](https://gorm.io): ORM库
- [golang-migrate](https://github.com/golang-migrate/migrate): 数据库迁移工具
- [JWT](https://github.com/golang-jwt/jwt): 认证