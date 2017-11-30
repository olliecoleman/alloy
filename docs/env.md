# Environment Variables file example

Placing a `.env` file at the root with the following schema:

```
export PORT=1212
export ENVIRONMENT=development
export HOST=localhost
export DATABASE_URL="host=127.0.0.1 port=5432 user=<user> password=<password> dbname=<db_name> sslmode=disable"
export DATABASE_URL="host=127.0.0.1 port=5432 user=<user> password=<password> dbname=<db_name> sslmode=disable"
export MIGRATIONS_DIR=app/migrations
export CSRF_KEY=XXX
export HASH_KEY=XXX
export BLOCK_KEY=XXX
export SMTP_HOST=<smtp_host>
export SMTP_PORT=<smtp_port>
export SMTP_USERNAME=<your_smtp_username>
export SMTP_PASSWORD=<your_smtp_password>
export MAILER_FROM=<some_email>
export MAILER_HOST=http://localhost:1212
export ASSET_URL=http://localhost:1212
```
