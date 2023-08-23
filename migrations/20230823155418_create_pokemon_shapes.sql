-- +goose Up
-- +goose StatementBegin
create table if not exists pokemon_shapes (
    id serial,
    identifier varchar(100)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table pokemon_shapes;
-- +goose StatementEnd
