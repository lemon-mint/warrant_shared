# Local development

## Set up Go

Warrant is written in Go. Prior to cloning the repo and making any code changes, please ensure that your local Go environment is set up. Refer to the appropriate instructions for your platform [here](https://go.dev/).

## Fork & clone repository

We follow GitHub's fork & pull request model. If you're looking to make code changes, it's easier to do so on your own fork and then contribute pull requests back to the Warrant repo. You can create your own fork of the repo [here](https://github.com/warrant-dev/warrant/fork).

If you'd just like to checkout the source to build & run, you can clone the repo directly:

```shell
git clone git@github.com:warrant-dev/warrant.git
```

Note: It's recommended you clone the repository into a directory relative to your `GOPATH` (e.g. `$GOPATH/src/github.com/warrant-dev`)

## Server configuration

Warrant requires certain configuration variables to be set via either a `warrant.yaml` config file or via environment variables. There is a set of common variables as well as datastore and eventstore-specific configuration.

### Common variables

| Variable | Description | Required? | Default | YAML | ENV VAR |
| -------- | ----------- | --------- | ------- | ---- | ------- |
| `port` | Port where the server runs. | no | 8000 | `port: VALUE` | `WARRANT_PORT=VALUE` |
| `logLevel` | Log level (e.g. Debug, Info etc.) for the server. Warrant uses zerolog, valid log levels are defined [here](https://github.com/rs/zerolog#leveled-logging). | no | 0 | `logLevel: VALUE` | `WARRANT_LOGLEVEL=VALUE` |
| `enableAccessLog` | Determines whether the built-in request logger is enabled or not. | no | true | `enableAccessLog: VALUE` | `WARRANT_ENABLEACCESSLOG=VALUE` |
| `autoMigrate` | If set to `true`, the server will apply datastore and eventstore migrations before starting up. | no | false | `autoMigrate: VALUE` | `WARRANT_AUTOMIGRATE=VALUE` |
| `authentication.apiKey` | The unique API key that all clients must pass to the server via the `Authorization: ApiKey VALUE` header | yes | - | `authentication:`<br>&emsp;`apiKey: VALUE` | `WARRANT_AUTHENTICATION_APIKEY=VALUE` |

## Set up datastore & eventstore

Warrant is a stateful service that runs with an accompanying `datastore` and `eventstore` (for tracking resource & access events). Currently, `MySQL`, `PostgreSQL` and `SQLite` (file and in-memory) are supported. Refer to these guides to set up your desired database(s):

- [MySQL](/migrations/datastore/mysql/README.md)
- [PostgreSQL](/migrations/datastore/postgres/README.md)
- [SQLite](/migrations/datastore/sqlite/README.md)

Note: It's possible to use different dbs for the `datastore` and `eventstore` (e.g. mysql for datastore and sqlite for eventstore) but we recommend using the same type of db during development for simplicity.

Here is an example of a full server config using `mysql` for both the datastore and eventstore:

### Sample `warrant.yaml` config (place file in same dir as server binary)

```yaml
port: 8000
logLevel: 1
enableAccessLog: true
autoMigrate: true
authentication:
    apiKey: your_api_key
datastore:
  mysql:
    username: replace_with_username
    password: replace_with_password
    hostname: replace_with_hostname
    database: warrant
eventstore:
  synchronizeEvents: false
  mysql:
    username: replace_with_username
    password: replace_with_password
    hostname: replace_with_hostname
    database: warrantEvents
```

### Sample environment variables config

```shell
export WARRANT_PORT=8000
export WARRANT_LOGLEVEL=1
export WARRANT_ENABLEACCESSLOG=true
export WARRANT_AUTOMIGRATE=true
export WARRANT_AUTHENTICATION_APIKEY="replace_with_api_key"
export WARRANT_DATASTORE_MYSQL_USERNAME="replace_with_username"
export WARRANT_DATASTORE_MYSQL_PASSWORD="replace_with_password"
export WARRANT_DATASTORE_MYSQL_HOSTNAME="replace_with_hostname"
export WARRANT_DATASTORE_MYSQL_DATABASE="warrant"
export WARRANT_EVENTSTORE_SYNCHRONIZEEVENTS=false
export WARRANT_EVENTSTORE_MYSQL_USERNAME="replace_with_username"
export WARRANT_EVENTSTORE_MYSQL_PASSWORD="replace_with_password"
export WARRANT_EVENTSTORE_MYSQL_HOSTNAME="replace_with_hostname"
export WARRANT_EVENTSTORE_MYSQL_DATABASE="warrantEvents"
```

## Build binary & start server

After the datastore, eventstore and configuration are set, build & start the server:

```shell
cd cmd/warrant
make dev
./bin/warrant
```

## Make requests

Once the server is running, you can make API requests using curl, any of the [Warrant SDKs](/README.md#sdks), or your favorite API client:

```shell
curl -g "http://localhost:port/v1/object-types" -H "Authorization: ApiKey YOUR_KEY"
```

# Running tests

## Unit tests

```shell
go test -v ./...
```

## End-to-end API tests

The Warrant repo contains a suite of e2e tests that test various combinations of API requests. These tests are defined in json files within the `tests/` dir and are executed using [APIRunner](https://github.com/warrant-dev/apirunner). These tests can be run locally:

### Install APIRunner

```shell
go install github.com/warrant-dev/apirunner/cmd/apirunner@latest
```

### Define test configuration

APIRunner tests run based on a simple config file that you need to create in the `tests/` directory:

```shell
touch tests/apirunner.conf
```

Add the following to your `tests/apirunner.conf` (replace with your server url and api key):

```json
{
    "baseUrl": "YOUR_SERVER_URL",
    "headers": {
        "Authorization" : "ApiKey YOUR_API_KEY"
    }
}
```

### Run tests

First, make sure your server is running:

```shell
./bin/warrant
```

In a separate shell, run the tests:

```shell
cd tests/
apirunner .
```