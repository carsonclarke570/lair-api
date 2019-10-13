# Lair API

API for Lair, a Dungeons and Dragons 5E campaign manager.

## Getting Started

### Prerequisites

Make sure you have Docker installed on your system and that the Docker daemon is running.

### Local Development

To build and test a local instance of the Lair API, simply run:

```
docker-compose run api
```

This initializes a local MySQL database using `sql/init.sql` and populates using `sql/populate.sql`

### API

Each model in Lair's API is subject to generic CRUD handlers. The behavior and usage of these handlers is as follows:

#### CREATE

```
POST www.example-host.com/api/users
```

Body:
```
{
    "first_name": "Example",
    "last_name": "User",
    "email": "user@example.com"
}
```

Create uses all the fields in the request's body (except for `ID`, `Created` and `Modified`) to create a new entry of that type in the database. It also generates a new ID for the model and sets the `Modified` and `Created` fields to `time.Now()`.

#### READ

```
GET www.example-host.com/api/users/{id}
```

Read returns the database entry at `{id}` as JSON in the response's body.

#### UPDATE

```
PUT www.example-host.com/api/users/{id}
```

Body:
```
{
    "last_name": "Admin",
}
```

Update finds the database entry at `{id}` and then updates each field in the request's body at that entry.

#### DELETE

```
DELETE www.example-host.com/api/users/{id}
```

Deletes the database entry at `{id}`

#### FILTER

```
GET www.example-host.com/api/users?last_name=Charles&order=asc&order_by=id&n=20&page=2
```

Filter returns a JSON list of database entries based on conditions passed as query parameters. Any field of a model can be used as a filter condition, in addition to the following:

| Parameter | Description | Default |
|-----------|-------------|---------|
| order | Determines ascending or descending order. Either `asc` or `des` | `asc` |
| order_by | Determines which field to order results by. | `id` |
| n | Determines how many results to display per page. | `20` |
| page | Determines which page of results to draw from. | `1` |

## Contributing
Feel free open a pull request and a Code Owner will get around to reviewing it. Don't hesitate to open a new issue or contact an author with any questions you may have.

## Authors

* **Carson Clarke-Magrab** - *Initial work* - [carsonclarke570](https://github.com/carsonclarke570)

## License

TO-DO: Add license

## Acknowledgments

* The UberEats app
