# Shopping list example

This example contains only one domain entity: an article.

## Prerequisites

* Docker and Docker Compose
* Go >= 1.22

## Configuration

Check out the dotfile `.env` at the project's root directory.

## Start the server

1. First start the database:

```sh
docker compose up -d
```

The schema will be created on the first start, the SQL script is `init_db.sql`.

2. Run the app's server:

```sh
go run ./main start-server
```

## API endpoints

The server starts on 0.0.0.1:8888.

* `GET /article`: List the saved articles.
* `POST /article`: Create a new article. Takes a JSON payload `{"name": "article name"}`.
* `GET /article/{id}`: Fetch the article ID `{id}`.
* `PUT /artcle/{id}`: Update the article ID `{id}`. Takes a JSON payload `{"name": "new article name"}`.

By default the endpoints return JSON responses (content type and response body). You can set a request header `Accept: application/xml`
to generate XML responses.
