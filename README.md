
# Documentation

A brief description of what this project does and who it's for


## Requirement

- Go version 1.17
- PostgreSQL 12.12


## Run Locally

Create project directory

```bash
  mkdir my-project
```

Clone the project

```bash
  git clone https://github.com/xcodeme21/go-crm-service.git
```

Go to the project directory

```bash
  cd my-project
```

Copy Environtment

```bash
  mv .env.sample .env || cp .env.sample .env
```

Setting Environtment

```bash
  setting [DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_TIMEZONE] according to your database settings
```

Install

```bash
  go get .
```

Running

```bash
  go run main.go
```


## API Reference

#### Get list source products

```http
  GET /list-source-products
```

#### Get list destination products

```http
  GET /list-destination-products
```

#### Auto Update destination products

```http
  GET /update-destination-products
```

