
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table batter_records (
    id  SERIAL PRIMARY KEY,
    player_id integer NOT NULL,
    belong_team_id smallint NOT NULL,
    year smallint NOT NULL,
    at_bats smallint NOT NULL,
    hit smallint NOT NULL,
    batting_average smallint NOT NULL,
    home_run smallint NOT NULL,
    rbi smallint NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp DEFAULT  NULL,
    foreign key (player_id) references players(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table batter_records;

