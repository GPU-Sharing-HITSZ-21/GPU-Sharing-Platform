USE `gpu_sharing_platform`;

INSERT INTO `test_instance` (`name`, `created_at`)
VALUES
    ('Instance 1', NOW()),
    ('Instance 2', NOW()),
    ('Instance 3', NOW()),
    ('Instance 4', NOW()),
    ('Instance 5', NOW());

INSERT INTO `users` (username, password, role, created_at)
VALUES ('admin_user', 'your_password_here', 'ADMIN', NOW());

INSERT INTO `ip_mapping` (public_ip, private_ip) VALUES
    ('110.40.176.8', '10.0.12.11'),
    ('124.223.53.29', '10.0.12.15');

INSERT INTO `pods` (pod_name, username, ssh_address, port_num) VALUES
('basePod','test','example',-1)