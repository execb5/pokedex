-- +goose Up
-- +goose StatementBegin
create table if not exists pokemon_colors (
    id serial primary key,
    identifier varchar(100)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table pokemon_colors;
-- +goose StatementEnd
