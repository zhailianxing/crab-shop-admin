# 环境配置使用指南

本项目使用环境配置文件来管理不同环境（开发、测试、生产）的配置。配置文件位于 `config` 目录下。

## 配置文件

项目包含以下环境配置文件：

- `config/local.env` - 本地开发环境配置（默认）
- `config/dev.env` - 开发/测试环境配置
- `config/prod.env` - 生产环境配置

## 配置项说明

每个环境配置文件包含以下主要配置项：

### 服务器配置

```
PORT=8080             # 服务器端口
ENV=development       # 环境名称
GIN_MODE=debug        # Gin模式（debug或release）
```

### 数据库配置

```
DB_HOST=localhost     # 数据库主机
DB_PORT=3306          # 数据库端口
DB_USER=root          # 数据库用户
DB_PASSWORD=password  # 数据库密码
DB_NAME=crabshop      # 数据库名称
DB_CHARSET=utf8mb4    # 数据库字符集
```

### JWT配置

```
JWT_SECRET=secret-key # JWT密钥
JWT_EXPIRATION=24h    # JWT过期时间
```

## 启动应用

您可以使用命令行参数 `-env` 指定要使用的环境配置文件：

### 使用本地开发环境配置启动（默认）

```bash
go run main.go
```

或

```bash
./admin.exe
```

您也可以明确指定使用本地环境配置：

```bash
go run main.go -env local
```

### 使用开发/测试环境配置启动

```bash
go run main.go -env dev
```

或

```bash
./admin.exe -env dev
```

### 使用生产环境配置启动

```bash
go run main.go -env prod
```

或

```bash
./admin.exe -env prod
```

## 自定义环境配置

您可以通过以下步骤创建自定义环境配置：

1. 在 `config` 目录中创建新的环境配置文件，例如 `staging.env`
2. 复制现有配置文件的内容，并根据需要修改配置值
3. 使用 `-env staging` 参数启动应用

## 环境变量优先级

加载环境变量的优先级如下：

1. 命令行设置的环境变量
2. 操作系统环境变量
3. 指定的环境配置文件（如 `dev.env` 或 `prod.env`）
4. 默认环境配置文件（`local.env`）

这意味着您可以通过环境变量覆盖配置文件中的设置，例如：

```bash
PORT=3000 ./admin.exe -env prod
```

上述命令将使用生产环境配置，但会覆盖端口设置为3000。

## 配置文件加载顺序

当指定的环境配置文件不存在时，系统会按照以下顺序尝试加载其他配置文件：

1. `config/local.env`（默认配置文件）
2. `config/.env`（旧版默认配置文件）
3. 项目根目录下的 `local.env`
4. 项目根目录下的 `.env`

如果所有文件都不存在，系统将只使用环境变量中的配置。 