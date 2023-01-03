# Notes

## Database CRUD Operations

DATABASE/SQL:

- Very fast & straightforward
- Manual mapping SQL fields to variables
- Easy to make mistakes, not caught until runtime

GORM:

- CRUD functions already implemented, very short production code
- Must learn to write queries using gorm's function
- Run slowly on high load

SQLX:

- Quite fast & easy to use
- Fields mappings via query text & struct tags
- FAilure won't occur until runtime

SQLC:

- Very fast & easy to use
- Automatic code generation
- Catch SQL query errors before generating codes
- Full supports Postgres, MySQL is experimental(? check)

## Validation Package

- it might give an error if the given query params value is zero, so make research about this.
