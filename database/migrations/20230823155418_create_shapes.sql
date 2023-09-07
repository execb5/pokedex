-- +goose Up
-- +goose StatementBegin
create table if not exists shapes (
    id serial primary key,
    identifier varchar(100)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table shapes;
-- +goose StatementEnd
