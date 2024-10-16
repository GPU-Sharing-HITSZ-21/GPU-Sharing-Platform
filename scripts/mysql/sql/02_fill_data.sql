USE `gpu_sharing_platform`;

INSERT INTO `test_instance` (`name`, `created_at`)
VALUES
    ('Instance 1', NOW()),
    ('Instance 2', NOW()),
    ('Instance 3', NOW()),
    ('Instance 4', NOW()),
    ('Instance 5', NOW());

INSERT INTO `user` (username, password, role, created_at)
VALUES ('admin_user', 'your_password_here', 'ADMIN', NOW());