-- 创建数据库
CREATE DATABASE IF NOT EXISTS gin_grom CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 使用数据库
USE gin_grom;

-- 创建管理员用户（密码：admin123）
INSERT INTO users (username, email, password, role, status, created_at, updated_at) 
VALUES (
    'admin', 
    'admin@example.com', 
    '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 
    'admin', 
    1, 
    NOW(), 
    NOW()
) ON DUPLICATE KEY UPDATE id=id;