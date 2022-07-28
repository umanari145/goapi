
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table users (
    id  SERIAL PRIMARY KEY,
    user_name varchar(50) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp DEFAULT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table users;
