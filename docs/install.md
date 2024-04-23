## Install


### Docker Setup

See the README in the [.docker-setup](https://github.com/dawitalemu4/postwoman/tree/main/.docker-setup) folder for the docker setup guide.


### Local Setup

To locally run postwoman, you need to have Go, PostgreSQL, Bash, and cURL installed on your machine.

1. Download the ZIP of this repo or clone the repository
```bash
git clone https://github.com/dawitalemu4/postwoman.git
```

2. Install the dependencies
```bash
go mod tidy
```

3. Rename the `.env.example` file to `.env` and use your own values (or you can just use the provided values)

4. Start the PostgreSQL server
```bash
pg_ctl -D "C:\Program Files\PostgreSQL\16\data" start # default postgres path on windows
```
or
```bash
pg_ctl -D /usr/local/var/postgres start # default postgres path on mac
```

5. Run the server (I prefer air for hot reload)
```bash
go run server.go
```
or
```bash
air
```

6. Open your browser and navigate to `localhost:YOURPORT`


Download links: [Go](https://go.dev/doc/install), [PostgreSQL](https://www.postgresql.org/download/), [Bash](https://git-scm.com/downloads) (I prefer git bash), [cURL](https://curl.se/download.html).

Contribute to add your OS! Read [postwoman.dev/contributing](https://postwoman.dev/contributing) or [contributing.md](https://github.com/dawitalemu4/postwoman/blob/main/docs/contributing.md).
