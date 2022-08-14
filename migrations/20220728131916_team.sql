
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table teams (
    id  SERIAL PRIMARY KEY,
    team_name varchar(20) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp DEFAULT  NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table teams;
