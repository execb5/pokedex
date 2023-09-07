-- +goose Up
-- +goose StatementBegin
create table if not exists colors (
    id serial primary key,
    identifier varchar(100)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table colors;
-- +goose StatementEnd
