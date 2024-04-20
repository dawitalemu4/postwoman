<h1 align="center">
    <img src="https://github.com/dawitalemu4/postwoman/assets/106638403/7555cd2c-3cf0-42fa-9420-90c35502a897" alt="postwoman icon" style="width: 50px; height: 50px;">
    <p>postwoman</p>
</h1>

Postwoman is a self-hosted dev tool that achieves what postman does but with my personal preferences in its functionality and design. Made with Go (Echo), HTMX, and cURL.

## Differences from Postman

- Never touch your mouse again to test your API (if you want to)
- No need to sign up or log in and give your data to a third party
- But keep your history and build profiles to save favorite requests (and autofill the form with the saved request!)
- Less cluttered UI and makes the most important things always accessible
- And more...

## Features

video demos ...

## Installation

To locally run postwoman, you need to have Go, PostgreSQL, Bash, and cURL installed on your machine.

1. Download the ZIP of this repo or clone the repository
```bash
git clone https://github.com/dawitalemu4/postwoman.git
```

2. Install the dependencies
```bash
go mod tidy
```

3. Rename the `.env.example` file to `.env` and use your own values

4. Start the PostgreSQL server (the default windows path is "C:\Program Files\PostgreSQL\16\data")
```bash
pg_ctl -D "YOURPATH" start
```

4. Run the server (I prefer air for hot reload)
```bash
go run server.go
```
or
```bash
air
```

5. Open your browser and navigate to `localhost:YOURPORT`

Download links: [Go](https://go.dev/doc/install), [PostgreSQL](https://www.postgresql.org/download/), [Bash](https://git-scm.com/downloads) (I prefer git bash), [cURL](https://curl.se/download.html).

See the README in the [.docker-setup](https://github.com/dawitalemu4/postwoman/tree/main/.docker-setup) folder for the docker installation guide.

## Shorter Startup

Check out my [startup script](https://github.com/dawitalemu4/postwoman/blob/main/startup.sh) to easily start up postwoman locally from a shortcut on your taskbar.

## Contributing

I'm open to contributions and suggestions, but fork this project if there are any crazy big changes you want to make that go against the [things I want to keep](https://github.com/dawitalemu4/postwoman/tree/main/docs/contributing.md).

Follow the template in the [contributing.md](https://github.com/dawitalemu4/postwoman/tree/main/docs/contributing.md) if you create a pull request or an issue.

## FAQ

## License