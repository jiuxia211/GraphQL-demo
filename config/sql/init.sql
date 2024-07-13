-- 创建用户表
CREATE TABLE user (
    `id` bigint AUTO_INCREMENT PRIMARY KEY,
    `username`   varchar(255) NOT NULL UNIQUE,
    `password`   varchar(255) NOT NULL, -- 存储哈希后的密码
    `email`      varchar(255) NOT NULL UNIQUE,
    `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
	`updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
	`deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间'
)ENGINE = InnoDB AUTO_INCREMENT = 1000 CHARSET = utf8mb4;

-- 创建备忘录表
CREATE TABLE task (
    `id` bigint AUTO_INCREMENT PRIMARY KEY,
    `uid` bigint,  
    `title` varchar(255),
    `content` longtext ,
    `status` tinyint(1),
    `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
	`updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
	`deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间'
)ENGINE = InnoDB AUTO_INCREMENT = 1000 CHARSET = utf8mb4;
