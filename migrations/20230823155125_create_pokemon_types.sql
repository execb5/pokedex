-- +goose Up
-- +goose StatementBegin
create table if not exists pokemon_types (
    id serial primary key,
    type_id int not null references types,
    slot int not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table pokemon_types;
-- +goose StatementEnd
