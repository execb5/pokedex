-- +goose Up
-- +goose StatementBegin
create table if not exists pokemon_types (
    pokemon_id serial references pokemon,
    type_id int not null references types,
    slot int not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table pokemon_types;
-- +goose StatementEnd
