-- +goose Up

CREATE TABLE todos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    valid_till TIMESTAMP,
    completed BOOLEAN,
    completed_at TIMESTAMP
);

-- +goose Down

DROP TABLE todos;