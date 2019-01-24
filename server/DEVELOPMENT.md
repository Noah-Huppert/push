# Development
Development instructions.

# Table Of Contents
- [Overview](#overview)
- [Dependencies](#dependencies)
- [Configuration](#configuration)
- [Local Database](#local-database)
- [Run](#run)

# Overview
When you first setup the repository complete steps in the
[Dependencies](#dependencies) section.

To run the server complete steps in the [Run](#run) section.

# Dependencies
[Dep](https://golang.github.io/dep/) is used to manage dependencies.  

Run the following command when you clone the repository or if you add a 
new dependency:

```
dep ensure
```

# Configuration
Set configuration environment variables:

- `PUSH_SERVER_HTTP_PORT`
- `PUSH_SERVER_DB_NAME`
- `PUSH_SERVER_DB_USER`

See the [`config/config.go`](./config/config.go) file for documentation on 
these values.

#Local Database
To start a local database Docker is required.

Run:

```
make db
```

# Run
Complete steps in [Configuration](#configuration) and 
[Local Database](#local-database).

Then run:

```
make
```
