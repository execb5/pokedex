# Pokedex

This is a toy project I've been working on to teach myself the
[go](https://go.dev/) programming language.

## Data and Database

All data and database structure are from the
[pokedata](https://github.com/ninjarobot/pokedata) repo, but I ended up
adding keys and changing the names a little bit. Also had to change the
order of rows on the pokemon csv to make it easier to be inserted since
now it has a self referencing key.

### Migrations

Migrations are managed using [goose](https://github.com/pressly/goose).

#### Run all pending migrations
```bash
goose -dir database/migrations postgres "host=localhost user=pokedex password=pokedex dbname=pokedex port=5432 sslmode=disable" up
```

#### Reset the db running the down of every migration
```bash
goose -dir database/migrations postgres "host=localhost user=pokedex password=pokedex dbname=pokedex port=5432 sslmode=disable" reset
```

#### Down the last applied migration
```bash
goose -dir database/migrations postgres "host=localhost user=pokedex password=pokedex dbname=pokedex port=5432 sslmode=disable" down
```

## Air

I've been using [Air](https://github.com/cosmtrek/air) for live
reloading. The only change I've made to the default configuration is the
executable name since it would've used  `main` otherwise.
