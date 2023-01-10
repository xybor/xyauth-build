# Get started

If you don't want to setup everything, let start with [Docker](#docker).

## Prerequisites

- Ubuntu 18.04
- Golang 1.18
- Postgres 12.12
- Make
- OpenSSL (for generate certificate)

## Setup database

This [article](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-postgresql-on-ubuntu-18-04) will guide you to setup the postgres database.

Then setup your new user and password [here](https://ubiq.co/database-blog/create-user-postgresql/).

After all, you need to determine the user, password, and database name to access to your database.

## Generate server certificate

Following this [guide](https://devopscube.com/create-self-signed-certificates-openssl/) to create your own certificate. The output is the private key (`server.key`) and public key (`server.crt`).

The following command will help you to generate a temporary certifcate:

```bash
make cert-gen
```

## Setup configurations

Setup your private key, public key, and other configurations in [config file](./configs/default.ini).

## Setup Environment Variables

Please refer [.env.example](./.env.example) to setup the environment variables.

```bash
export key=value
```

Specially, if the value of `general.environment` in [config file](./configs/default.ini) is `dev`, you can create `.env` file from [.env.example](./.env.example) instead of using `export`.

## Run the server

```bash
make run
```

# Docker

## Prerequisites

- Docker
- Docker compose
- Make
- (Optional) OpenSSL

## Generate server certificate

```bash
make cert-gen
```

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

## Stop docker-compose

```bash
make docker-stop
```

## Clean docker containers and images

```bash
make docker-clean
```
