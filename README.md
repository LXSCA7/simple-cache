# Simple cache

A basic caching implementation in Go, designed to illustrate how data can be stored and retrieved from a temporary location to reduce the load on a primary data source, such as a SQL database. This project focuses on demonstrating and learning the fundamental principles of caching using a local file system.

## How to run:

1. Clone the repository
```bash
git clone https://github.com/LXSCA7/simple-cache
cd simple-cache
```

2. Get dependencies
```bash
go mod tidy
```

3. Run the code
First time (for populate the database):
```bash
go run . insert
```

Subsequent runs:
```bash
go run .
``` 
