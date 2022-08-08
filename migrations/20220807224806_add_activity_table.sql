-- +goose Up
-- +goose StatementBegin
CREATE TABLE activities (
    id int AUTO_INCREMENT NOT NULL,
    check_in_id int NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id),
    FOREIGN KEY (check_in_id)
        REFERENCES check_ins(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE activities;
-- +goose StatementEnd
