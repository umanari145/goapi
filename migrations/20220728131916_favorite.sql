
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table favorites (
    id  SERIAL PRIMARY KEY,
    user_id integer NOT NULL,
    video_id varchar(255) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp DEFAULT  NULL,
    foreign key (user_id) references users(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table favorites;
