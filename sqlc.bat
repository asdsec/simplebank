:: This file has been written for windows user.
:: Because sqlc tool does not work properly with Git Bash.
:: In the project path, run `sqlc.bat` via cmd.

docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate