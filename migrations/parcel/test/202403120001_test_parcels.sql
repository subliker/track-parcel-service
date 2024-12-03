-- +goose Up
-- Добавление тестовых данных для таблицы parcels
INSERT INTO parcels (track_number, name, manager_id, recipient, arrival_address, forecast_date, description)
VALUES
    ('00000000001', 'Parcel 1', 101, 'John Doe', '123 Main St, Springfield', '2024-12-05 10:00:00', 'Electronics item'),
    ('00000000002', 'Parcel 2', 102, 'Jane Smith', '456 Elm St, Metropolis', '2024-12-06 14:00:00', 'Clothing package'),
    ('00000000003', 'Parcel 3', 103, 'Alice Johnson', '789 Oak St, Gotham', '2024-12-07 09:30:00', 'Books and stationery'),
    ('00000000004', 'Parcel 4', 104, 'Bob Brown', '101 Maple Ave, Star City', '2024-12-05 16:45:00', 'Sports equipment'),
    ('00000000005', 'Parcel 5', 105, 'Charlie White', '202 Pine Rd, Central City', '2024-12-08 08:15:00', 'Gadget accessories'),
    ('00000000006', 'Parcel 6', 106, 'Diana Prince', '303 Birch Blvd, Coast City', '2024-12-09 11:20:00', 'Cosmetics and toiletries'),
    ('00000000007', 'Parcel 7', 107, 'Clark Kent', '404 Cedar Ct, Smallville', '2024-12-10 13:00:00', 'Food items'),
    ('00000000008', 'Parcel 8', 108, 'Bruce Wayne', '505 Willow Ln, Gotham', '2024-12-05 15:00:00', 'Furniture pieces'),
    ('00000000009', 'Parcel 9', 109, 'Barry Allen', '606 Redwood St, Central City', '2024-12-06 10:45:00', 'Electronic components'),
    ('00000000010', 'Parcel 10', 110, 'Hal Jordan', '707 Spruce Dr, Coast City', '2024-12-07 12:30:00', 'Home decor items'),
    ('00000000011', 'Parcel 11', 101, 'John Doe', '123 Main St, Springfield', '2024-12-05 10:00:00', 'Electronics item'),
    ('00000000012', 'Parcel 12', 102, 'Jane Smith', '456 Elm St, Metropolis', '2024-12-06 14:00:00', 'Clothing package'),
    ('00000000013', 'Parcel 13', 103, 'Alice Johnson', '789 Oak St, Gotham', '2024-12-07 09:30:00', 'Books and stationery'),
    ('00000000014', 'Parcel 14', 104, 'Bob Brown', '101 Maple Ave, Star City', '2024-12-05 16:45:00', 'Sports equipment'),
    ('00000000015', 'Parcel 15', 105, 'Charlie White', '202 Pine Rd, Central City', '2024-12-08 08:15:00', 'Gadget accessories'),
    ('00000000016', 'Parcel 16', 106, 'Diana Prince', '303 Birch Blvd, Coast City', '2024-12-09 11:20:00', 'Cosmetics and toiletries'),
    ('00000000017', 'Parcel 17', 107, 'Clark Kent', '404 Cedar Ct, Smallville', '2024-12-10 13:00:00', 'Food items'),
    ('00000000018', 'Parcel 18', 108, 'Bruce Wayne', '505 Willow Ln, Gotham', '2024-12-05 15:00:00', 'Furniture pieces'),
    ('00000000019', 'Parcel 19', 109, 'Barry Allen', '606 Redwood St, Central City', '2024-12-06 10:45:00', 'Electronic components'),
    ('00000000020', 'Parcel 20', 110, 'Hal Jordan', '707 Spruce Dr, Coast City', '2024-12-07 12:30:00', 'Home decor items');

-- +goose Down
-- Удаление тестовых данных из таблицы parcels
DELETE FROM parcels
WHERE track_number IN (
    '00000000001', '00000000002', '00000000003', '00000000004', '00000000005',
    '00000000006', '00000000007', '00000000008', '00000000009', '00000000010',
    '00000000011', '00000000012', '00000000013', '00000000014', '00000000015',
    '00000000016', '00000000017', '00000000018', '00000000019', '00000000020'
);
