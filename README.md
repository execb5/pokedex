# Pokedex

All data is from [pokedata](https://github.com/ninjarobot/pokedata).

Migrations are managed using [goose](https://github.com/pressly/goose).

```bash
# Run all pending migrations
goose -dir migrations postgres "host=localhost user=pokedex password=pokedex dbname=pokedex port=5432 sslmode=disable" up

# Reset the db running the down of every migration
goose -dir migrations postgres "host=localhost user=pokedex password=pokedex dbname=pokedex port=5432 sslmode=disable" reset

# Down the last applied migration
goose -dir migrations postgres "host=localhost user=pokedex password=pokedex dbname=pokedex port=5432 sslmode=disable" down
```
