App entry point is in `cmd/todolist/main.go`

Following .env variables are necessary:
- db_mode ("postgres" or "inmem")
- When using postgres
    - host
    - port
    - user
    - password
    - database
    - sslmode
