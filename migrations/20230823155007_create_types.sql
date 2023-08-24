-- +goose Up
-- +goose StatementBegin
create table if not exists types (
    id serial primary key,
    identifier varchar(100) not null,
    generation_id int not null,
    damage_class_id int
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table types;
-- +goose StatementEnd
