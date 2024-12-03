-- +goose Up
-- Добавление тестовых данных для таблицы users
INSERT INTO users (telegram_id, full_name, email, phone_number)
VALUES
    (201, 'John Doe', 'johndoe@example.com', '+1234567890'),
    (202, 'Jane Smith', 'janesmith@example.com', '+1234567891'),
    (203, 'Alice Johnson', 'alicejohnson@example.com', '+1234567892'),
    (204, 'Bob Brown', 'bobbrown@example.com', '+1234567893'),
    (205, 'Charlie White', 'charliewhite@example.com', '+1234567894'),
    (206, 'Diana Prince', 'dianaprince@example.com', '+1234567895'),
    (207, 'Clark Kent', 'clarkkent@example.com', '+1234567896'),
    (208, 'Bruce Wayne', 'brucewayne@example.com', '+1234567897'),
    (209, 'Barry Allen', 'barryallen@example.com', '+1234567898'),
    (210, 'Hal Jordan', 'haljordan@example.com', '+1234567899'),
    (211, 'Oliver Queen', 'oliverqueen@example.com', '+1234567810'),
    (212, 'Arthur Curry', 'arthurcurry@example.com', '+1234567811'),
    (213, 'Victor Stone', 'victorstone@example.com', '+1234567812'),
    (214, 'Billy Batson', 'billybatson@example.com', '+1234567813'),
    (215, 'Selina Kyle', 'selinakyle@example.com', '+1234567814'),
    (216, 'Harley Quinn', 'harleyquinn@example.com', '+1234567815'),
    (217, 'Lex Luthor', 'lexluthor@example.com', '+1234567816'),
    (218, 'Lois Lane', 'loislane@example.com', '+1234567817'),
    (219, 'Jimmy Olsen', 'jimmyolsen@example.com', '+1234567818'),
    (220, 'Alfred Pennyworth', 'alfredpennyworth@example.com', '+1234567819');

-- +goose Down
-- Удаление тестовых данных из таблицы users
DELETE FROM users
WHERE telegram_id BETWEEN 201 AND 220;
