-- +goose Up
-- +goose StatementBegin
CREATE TABLE check_outs (
    id int AUTO_INCREMENT NOT NULL,
    check_in_id int NOT NULL,
    date_check_out TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY (check_in_id)
        REFERENCES check_ins(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE check_outs;
-- +goose StatementEnd
