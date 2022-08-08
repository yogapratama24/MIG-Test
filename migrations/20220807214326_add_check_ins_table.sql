-- +goose Up
CREATE TABLE check_ins (
    id int AUTO_INCREMENT NOT NULL,
    date_check_in TIMESTAMP NOT NULL,
    user_id int NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);
-- +goose StatementBegin
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE check_ins;
-- +goose StatementEnd
