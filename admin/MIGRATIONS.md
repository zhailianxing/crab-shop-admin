# 数据库迁移指南

本项目使用 SQL 迁移文件管理数据库结构，不再使用 GORM 的自动迁移功能。这样可以更好地控制数据库变更，并确保所有环境（开发、测试、生产）中的数据库结构保持一致。

## 迁移文件结构

迁移文件存放在 `migrations` 目录下，按照以下命名规则：

```
<version>_<description>.<direction>.sql
```

例如：
- `000001_create_users_table.up.sql` - 创建用户表的迁移
- `000001_create_users_table.down.sql` - 回滚用户表创建的迁移

每个迁移都有两个文件：
- `.up.sql` - 向前迁移的 SQL 语句（创建表、添加列等）
- `.down.sql` - 向后迁移的 SQL 语句（删除表、删除列等）

## 创建新的迁移

### 1. 创建新表

如果需要创建新表，请按照以下步骤操作：

1. 在 `migrations` 目录中创建两个新文件：
   - `00000X_create_table_name.up.sql`
   - `00000X_create_table_name.down.sql`

2. 在 `.up.sql` 文件中编写创建表的 SQL 语句：
   ```sql
   CREATE TABLE IF NOT EXISTS `table_name` (
     `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
     `name` varchar(100) NOT NULL,
     `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
     `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
     `deleted_at` datetime(3) DEFAULT NULL,
     PRIMARY KEY (`id`),
     KEY `idx_table_name_deleted_at` (`deleted_at`)
   ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
   ```

3. 在 `.down.sql` 文件中编写回滚 SQL 语句：
   ```sql
   DROP TABLE IF EXISTS `table_name`;
   ```

### 2. 修改现有表结构

如果需要修改现有表（如添加列、修改列、添加索引等），请按照以下步骤操作：

1. 在 `migrations` 目录中创建两个新文件：
   - `00000X_alter_table_name.up.sql`
   - `00000X_alter_table_name.down.sql`

2. 在 `.up.sql` 文件中编写修改表的 SQL 语句：
   ```sql
   ALTER TABLE `table_name` 
   ADD COLUMN `new_column` varchar(255) DEFAULT NULL,
   ADD INDEX `idx_table_name_new_column` (`new_column`);
   ```

3. 在 `.down.sql` 文件中编写回滚 SQL 语句：
   ```sql
   ALTER TABLE `table_name`
   DROP COLUMN `new_column`;
   ```

## 更新模型

在创建或修改数据库表后，请确保更新对应的 Go 模型定义，使其与数据库结构保持一致：

```go
// internal/models/model_name.go
type ModelName struct {
    ID        uint      `gorm:"primarykey" json:"id"`
    Name      string    `gorm:"size:100;not null" json:"name"`
    NewColumn string    `gorm:"size:255" json:"new_column"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt *time.Time `gorm:"index" json:"-"`
}
```

## 应用迁移

使用以下命令应用迁移（向前）：

```bash
go run cmd/migrate/main.go -direction=up
```

使用以下命令回滚迁移（向后）：

```bash
go run cmd/migrate/main.go -direction=down
```

## 迁移注意事项

1. **保持向后兼容**：在添加新列时，尽量使用 `DEFAULT NULL` 或提供合适的默认值，以避免对现有数据的影响。

2. **小步迁移**：每次迁移尽量只做一个小的变更，这样更容易回滚和调试。

3. **测试迁移**：在应用到生产环境之前，先在开发或测试环境中测试迁移。

4. **备份数据**：在执行重要的迁移之前，务必备份数据库。

5. **维护版本序列**：确保迁移文件的版本号是连续的，并且反映了应用的顺序。 