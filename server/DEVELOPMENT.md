# Development
Development instructions.

# Table Of Contents
- [Overview](#overview)
- [Dependencies](#dependencies)
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

#Local Database
To start a local database Docker is required.

Run:

```
make db
```

# Run
Complete steps in the [Local Database](#local-database) section.

Then run:

```
make
```
