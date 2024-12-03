-- +goose Up
-- Добавление тестовых данных для таблицы managers
INSERT INTO managers (telegram_id, full_name, phone_number, email, company, api_token)
VALUES
    (101, 'John Manager', '+1234500001', 'john.manager@example.com', 'Acme Corp', 'API_TOKEN_101'),
    (102, 'Jane Manager', '+1234500002', 'jane.manager@example.com', 'Tech Solutions', 'API_TOKEN_102'),
    (103, 'Alice Manager', '+1234500003', 'alice.manager@example.com', 'Logistics Inc.', 'API_TOKEN_103'),
    (104, 'Bob Manager', '+1234500004', 'bob.manager@example.com', 'Global Trade', 'API_TOKEN_104'),
    (105, 'Charlie Manager', '+1234500005', 'charlie.manager@example.com', 'Fast Delivery', 'API_TOKEN_105'),
    (106, 'Diana Manager', '+1234500006', 'diana.manager@example.com', 'Express Corp', 'API_TOKEN_106'),
    (107, 'Clark Manager', '+1234500007', 'clark.manager@example.com', 'Super Logistics', 'API_TOKEN_107'),
    (108, 'Bruce Manager', '+1234500008', 'bruce.manager@example.com', 'Wayne Enterprises', 'API_TOKEN_108'),
    (109, 'Barry Manager', '+1234500009', 'barry.manager@example.com', 'Speedy Shipments', 'API_TOKEN_109'),
    (110, 'Hal Manager', '+1234500010', 'hal.manager@example.com', 'Star Freight', 'API_TOKEN_110');

-- +goose Down
-- Удаление тестовых данных из таблицы managers
DELETE FROM managers
WHERE telegram_id BETWEEN 101 AND 110;
