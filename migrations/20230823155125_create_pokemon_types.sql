-- +goose Up
-- +goose StatementBegin
create table if not exists pokemon_types (
    id serial,
    type_id int not null,
    slot int not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table pokemon_types;
-- +goose StatementEnd
