-- +goose Up
-- +goose StatementBegin
create table if not exists pokemon (
    id serial primary key,
    identifier varchar(100) not null,
    species_id int not null references pokemon_species,
    height int not null,
    weight int not null,
    base_experience int not null,
    "order" int not null,
    is_default boolean not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table pokemon;
-- +goose StatementEnd
