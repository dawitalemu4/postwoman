<h1 align="center">
    <img src="https://github.com/dawitalemu4/postwoman/assets/106638403/7555cd2c-3cf0-42fa-9420-90c35502a897" alt="postwoman icon" style="width: 50px; height: 50px;">
    <p>postwoman</p>
</h1>

Postwoman is a self-hosted dev tool that achieves what postman does but with my personal preferences in its functionality and design. Made with Go (Echo), HTMX, and cURL.


## Differences from Postman

- Use your keyboard to do everything if you want to (literally everything)
- Less cluttered UI and makes the most important things always accessible
- No need to sign up or log in and give your data to a third party
- But keep your request history and build profiles to save favorite requests (and autofill the form with the saved request!)
- And more...


## Features

### Make Request

<video height="600px" controls autoplay preload muted loop>
    <source src="https://github.com/dawitalemu4/postwoman/assets/106638403/b09125dd-565e-479f-be58-194261c2667f">
</video>

<p style="visibility:hidden">https://github.com/dawitalemu4/postwoman/assets/106638403/b09125dd-565e-479f-be58-194261c2667f</p>

> JSON and HTML responses are automatically formatted.


### History List

<video height="600px" controls autoplay preload muted loop>
    <source src="https://github.com/dawitalemu4/postwoman/assets/106638403/27d2a4b0-9890-43bd-9bfd-8f928445e011">
</video>

<p style="visibility:hidden">https://github.com/dawitalemu4/postwoman/assets/106638403/27d2a4b0-9890-43bd-9bfd-8f928445e011</p>


### Favorites List

<video height="600px" controls autoplay preload muted loop>
    <source src="https://github.com/dawitalemu4/postwoman/assets/106638403/9dd3a5b9-e647-48ec-932f-6b0bd84885bd">
</video>

<p style="visibility:hidden">https://github.com/dawitalemu4/postwoman/assets/106638403/9dd3a5b9-e647-48ec-932f-6b0bd84885bd</p>


### Fill Form from List

<video height="600px" controls autoplay preload muted loop>
    <source src="https://github.com/dawitalemu4/postwoman/assets/106638403/2b725e99-f3b3-4b6c-ba3c-33f6369b3e1d">
</video>

<p style="visibility:hidden">https://github.com/dawitalemu4/postwoman/assets/106638403/2b725e99-f3b3-4b6c-ba3c-33f6369b3e1d</p>


Go to [postwoman.dev/features](https://postwoman.dev/features) or [features.md](https://github.com/dawitalemu4/postwoman/blob/main/docs/features.md) to view all features.


## Installation

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

View [postwoman.dev/install](https://postwoman.dev/install) or [install.md](https://github.com/dawitalemu4/postwoman/blob/main/docs/install.md) for more detailed instructions for your OS.


## Startup Shortcuts

Check out my [startup script](https://github.com/dawitalemu4/postwoman/blob/main/startup.sh) to easily start up postwoman locally from a shortcut on your taskbar, or this [startup script](https://github.com/dawitalemu4/postwoman/tree/main/.docker-setup/startup.sh) if you are using docker.

Visit [postwoman.dev/shortcuts](https://postwoman.dev/shortcuts) or [shortcuts.md](https://github.com/dawitalemu4/postwoman/blob/main/docs/shortcuts.md) for demo videos and tutorials on how to make your own shortcut.


## Contributing

I'm open to contributions and suggestions, but fork this project if there are any crazy big changes you want to make that go against the [things I want to keep](https://github.com/dawitalemu4/postwoman/blob/main/docs/contributing.md#things-i-want-to-keep).

Follow the checklist in the [contributing.md](https://github.com/dawitalemu4/postwoman/tree/main/docs/contributing.md) if you create a pull request or an issue.


## FAQ

**Q:** **Why this UI style? Just use the terminal?**

**A:** I wanted to be able to see the history of my requests, favorite some requests when I'm working on a specific controller, and see the status of the request on the request's preview before I select it, which postman didn't have and I hated guessing and trying each one to see which one works.

**Q:** **Why self-hosted?**

**A:**  I don't want to pay for cloud resources for your convenience. Jokes aside, it's impossible to make API requests to an API that is locally hosted from postwoman if it was hosted on a deployed server (without having a user download something locally), and I plan on using this on APIs I'm building locally.

Leave a post in the [discussions](https://github.com/dawitalemu4/postwoman/discussions) if you have any questions.


## License

This project is licensed under the Creative Commons Attribution-NonCommercial 4.0 International Public License - see the [LICENSE.txt](https://github.com/dawitalemu4/postwoman/blob/main/LICENSE.txt).