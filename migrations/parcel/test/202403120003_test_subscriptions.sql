-- +goose Up
-- Добавление тестовых данных для таблицы subscriptions
INSERT INTO subscriptions (user_id, parcel_track_number)
VALUES
    (201, '00000000001'),
    (202, '00000000002'),
    (203, '00000000003'),
    (204, '00000000004'),
    (205, '00000000005'),
    (206, '00000000006'),
    (207, '00000000007'),
    (208, '00000000008'),
    (209, '00000000009'),
    (210, '00000000010'),
    (211, '00000000011'),
    (212, '00000000012'),
    (213, '00000000013'),
    (214, '00000000014'),
    (215, '00000000015'),
    (216, '00000000016'),
    (217, '00000000017'),
    (218, '00000000018'),
    (219, '00000000019'),
    (220, '00000000020');

-- +goose Down
-- Удаление тестовых данных из таблицы subscriptions
DELETE FROM subscriptions
WHERE parcel_track_number IN (
    '00000000001', '00000000002', '00000000003', '00000000004', '00000000005',
    '00000000006', '00000000007', '00000000008', '00000000009', '00000000010',
    '00000000011', '00000000012', '00000000013', '00000000014', '00000000015',
    '00000000016', '00000000017', '00000000018', '00000000019', '00000000020'
);