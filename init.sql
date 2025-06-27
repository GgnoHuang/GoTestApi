-- 创建simple_strings表
CREATE TABLE IF NOT EXISTS simple_strings (
    id SERIAL PRIMARY KEY,
    value TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 插入一些示例数据
INSERT INTO simple_strings (value) VALUES 
    ('Hello World'),
    ('Docker is awesome'),
    ('Go + PostgreSQL + Docker')
ON CONFLICT DO NOTHING; 