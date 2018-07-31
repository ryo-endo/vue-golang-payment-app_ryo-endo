DROP TABLE IF EXISTS items;

CREATE TABLE items (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255),
    description VARCHAR(255),
    amount INTEGER
);

INSERT INTO items (name, description, amount)
VALUES
    ('toy', 'test-toy', 2000);


INSERT INTO items (name, description, amount)
VALUES
    ('game', 'test-game', 1000);