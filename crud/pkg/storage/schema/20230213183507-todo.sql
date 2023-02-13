
-- +migrate Up
CREATE TABLE `campaign_analytics` (
    `id` varchar(50) NOT NULL,
    `name` text,
    `description` text,
    `priority` text,
    `duration` bigint,
    `status` varchar(10),
    `created_at` datetime,
    `updated_at` datetime,
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DELETE TABLE `todos`;