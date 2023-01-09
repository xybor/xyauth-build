# Get started

## Prerequisites

- Ubuntu 18.04
- Golang 1.18
- Postgres 12.12
- Make

## Setup Environment Variables

Please refer [.env.example](./.env.example) to setup the environment variables.

```bash
export key=value
```

Specially, if the value of `general.environment` in [config file](./configs/default.ini) is `dev`, you can create `.env` file from `.env.example` instead of using `export`.

## Build

```bash
make server
make database
```

## Run the server

```bash
make run
# or
./database migrate
./server
```

# Docker

## Prerequisites

- Docker
- Docker compose
- Make

Before running the following steps, please refer [Setup Environment Variables section](#setup-environment-variables).

## Generate dockerfile and docker compose

```bash
make docker-gen
```

## Build the image

```bash
make docker-build
```

## Start docker-compose

```bash
make docker-start
```

## Stop docker compose

```bash
make docker-stop
```

## Clean docker containers and images

```bash
make docker-clean
```
