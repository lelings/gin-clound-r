CREATE TABLE `users` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `uuid` varchar(50) NULL DEFAULT NULL COMMENT '用户uuid',
    `password` varchar(255) NOT NULL COMMENT '用户密码',
    `file_store_id` int(11) NULL DEFAULT NULL COMMENT '文件仓库ID',
    `username` varchar(20) NULL DEFAULT NULL COMMENT '用户名',
    `email` varchar(50) UNIQUE NULL DEFAULT NULL COMMENT '邮箱',
    `register_time` datetime NULL DEFAULT NULL COMMENT '注册时间',
    `image_path` varchar(255) NULL DEFAULT NULL COMMENT '用户头像',
    PRIMARY KEY(`id`) USING BTREE
);

CREATE TABLE my_files (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '文件id',
    `file_name` varchar(30) NULL DEFAULT NULL COMMENT '文件名',
    `file_hash` varchar(255) NULL DEFAULT NULL COMMENT '文件哈希值',
    `file_store_id` int(11) NULL DEFAULT NULL COMMENT '文件仓库id',
    `file_path` varchar(255) NULL DEFAULT '/' COMMENT '文件路径',
    `parent_folder_id` int(11) NULL DEFAULT NULL COMMENT '父文件夹id',
    `download_num` int(11) NULL DEFAULT 0 COMMENT '下载次数',
    `upload_time` varchar(50) NULL DEFAULT NULL COMMENT '上传时间',
    `type` int(11) NULL DEFAULT NULL COMMENT '文件类型',
    `size` int(11) NULL DEFAULT NULL COMMENT '文件大小',
    `size_str` varchar(50) NULL DEFAULT NULL COMMENT '文件大小单位',
    `postfix` varchar(50) NULL DEFAULT NULL COMMENT '文件后缀',
    PRIMARY KEY(`id`) USING BTREE
);

CREATE TABLE file_stores (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '仓库id',
    `user_id` int(11) NULL DEFAULT NULL COMMENT '用户id',
    `current_size` int(11) NULL DEFAULT NULL COMMENT '现在的大小',
    `max_size` int(11) NULL DEFAULT NULL COMMENT '最大大小',
    PRIMARY KEY(`id`) USING BTREE
);

CREATE TABLE file_folders (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '文件夹id',
    `file_folder_name` varchar(50) NULL DEFAULT NULL COMMENT '文件夹名',
    `file_store_id` int(11) NULL DEFAULT NULL COMMENT '所属仓库id',
    `parent_folder_id` int(11) NULL DEFAULT NULL COMMENT '父文件夹名',
    `time` varchar(50) NULL DEFAULT NULL COMMENT '创建时间',
    PRIMARY KEY(`id`) USING BTREE
);

CREATE TABLE share (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `file_id` int(11) NULL DEFAULT NULL COMMENT '分享的文件id',
    `code` varchar(20) NULL DEFAULT NULL COMMENT '提取码',
    `username` varchar(50) NULL DEFAULT NULL COMMENT '用户名',
    `hash` varchar(255) NULL DEFAULT NULL,
    PRIMARY KEY(`id`) USING BTREE
);