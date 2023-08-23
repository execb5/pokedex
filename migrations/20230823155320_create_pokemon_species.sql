-- +goose Up
-- +goose StatementBegin
create table if not exists pokemon_species (
    id serial,
    identifier varchar(100) not null,
    generation_id int,
    evolves_from_species_id int null,
    evolution_chain_id int,
    color_id int,
    shape_id int,
    habitat_id int,
    gender_rate int,
    capture_rate int,
    base_happiness int,
    is_baby boolean,
    hatch_counter int,
    has_gender_differences boolean,
    growth_rate_id int,
    forms_switchable boolean,
    "order" int,
    conquest_order int null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table pokemon_species;
-- +goose StatementEnd
