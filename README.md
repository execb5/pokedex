# Pokedex

## Data and Database

All data and database structure are from the
[pokedata](https://github.com/ninjarobot/pokedata) repo, but I ended up
adding keys and changing the names a little bit. Also had to change the
order of rows on the pokemon csv to make it easier to be inserted since
now it has a self referencing key.

### Migrations

Migrations are managed using [goose](https://github.com/pressly/goose).

```bash
# Run all pending migrations
goose -dir migrations postgres "host=localhost user=pokedex password=pokedex dbname=pokedex port=5432 sslmode=disable" up

# Reset the db running the down of every migration
goose -dir migrations postgres "host=localhost user=pokedex password=pokedex dbname=pokedex port=5432 sslmode=disable" reset

# Down the last applied migration
goose -dir migrations postgres "host=localhost user=pokedex password=pokedex dbname=pokedex port=5432 sslmode=disable" down
```
