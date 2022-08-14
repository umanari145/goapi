
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table pitcher_records (
    id  SERIAL PRIMARY KEY,
    player_id integer NOT NULL,
    belong_team_id smallint NOT NULL,
    year smallint NOT NULL,
    pitches smallint NOT NULL,
    win smallint NOT NULL,
    lose smallint NOT NULL,
    era smallint NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp DEFAULT  NULL,
    foreign key (player_id) references players(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table pitcher_records;
