-- +goose Up
-- Добавление тестовых данных для таблицы checkpoints
INSERT INTO checkpoints (time, place, description, parcel_track_number, parcel_status)
VALUES
    ('2024-12-01 10:00:00', 'Sorting Center A', 'Parcel received at sorting center', '00000000001', 'PENDING'),
    ('2024-12-01 12:00:00', 'Warehouse B', 'Parcel transferred to warehouse', '00000000001', 'IN_TRANSIT'),
    ('2024-12-02 08:00:00', 'Delivery Hub C', 'Out for delivery', '00000000001', 'DELIVERED'),
    ('2024-12-01 09:30:00', 'Sorting Center X', 'Parcel received at sorting center', '00000000002', 'PENDING'),
    ('2024-12-01 11:00:00', 'Warehouse Y', 'Parcel transferred to warehouse', '00000000002', 'IN_TRANSIT'),
    ('2024-12-02 07:30:00', 'Delivery Hub Z', 'Out for delivery', '00000000002', 'DELIVERED'),
    ('2024-12-01 11:15:00', 'Depot Q', 'Parcel processed at depot', '00000000003', 'PENDING'),
    ('2024-12-01 14:00:00', 'Warehouse R', 'Parcel transferred to warehouse', '00000000003', 'IN_TRANSIT'),
    ('2024-12-02 09:00:00', 'Delivery Hub S', 'Out for delivery', '00000000003', 'DELIVERED'),
    ('2024-12-01 08:45:00', 'Sorting Center M', 'Parcel received at sorting center', '00000000004', 'PENDING'),
    ('2024-12-01 10:30:00', 'Warehouse N', 'Parcel transferred to warehouse', '00000000004', 'IN_TRANSIT'),
    ('2024-12-02 08:15:00', 'Delivery Hub O', 'Out for delivery', '00000000004', 'DELIVERED'),
    ('2024-12-01 07:00:00', 'Sorting Center U', 'Parcel received at sorting center', '00000000005', 'PENDING'),
    ('2024-12-01 09:00:00', 'Warehouse V', 'Parcel transferred to warehouse', '00000000005', 'IN_TRANSIT'),
    ('2024-12-02 07:00:00', 'Delivery Hub W', 'Out for delivery', '00000000005', 'DELIVERED'),
    ('2024-12-01 06:30:00', 'Depot H', 'Parcel processed at depot', '00000000006', 'PENDING'),
    ('2024-12-01 08:15:00', 'Warehouse I', 'Parcel transferred to warehouse', '00000000006', 'IN_TRANSIT'),
    ('2024-12-02 06:45:00', 'Delivery Hub J', 'Out for delivery', '00000000006', 'DELIVERED'),
    ('2024-12-01 09:00:00', 'Sorting Center K', 'Parcel received at sorting center', '00000000007', 'PENDING'),
    ('2024-12-01 11:30:00', 'Warehouse L', 'Parcel transferred to warehouse', '00000000007', 'IN_TRANSIT');

-- +goose Down
-- Удаление тестовых данных из таблицы checkpoints
DELETE FROM checkpoints
WHERE parcel_track_number IN ('00000000001', '00000000002', '00000000003', '00000000004', '00000000005', '00000000006', '00000000007');