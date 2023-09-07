-- +goose Up
-- +goose StatementBegin
create table if not exists pokemon_types (
    pokemon_id int not null references pokemon,
    type_id int not null references types,
    slot int not null,

    constraint unique_pokemon_type unique (pokemon_id, type_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table pokemon_types;
-- +goose StatementEnd
