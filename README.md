# Traffic Control

Project for the subject Distributed Systems at the Cooperative State University Stuttgart.

## Production

### Requirements

#### Client

-   Docker (for Windows via Docker Desktop and WSL)
-   Go

#### Server

-   Docker

### Execution

You can find the detailed description to run

-   [server](https://github.com/Patr1ick/dhbw-traffic-control/blob/main/server/README.md)
-   [client](https://github.com/Patr1ick/dhbw-traffic-control/blob/main/client/README.md)

## Local Development Setup

To set up the local development environment, simply set up the corresponding docker compose.

```bash
# Build docker container
docker compose -f docker-compose-dev.yaml build
# Start docker container
docker compose -f docker-compose-dev.yaml up -d
```

This starts the Docker containers for the server and Apache Cassandra. Another container initialises the Apache Cassandra database with the respective KeySpace and table. In the local setup, Caddy is currently not considered. Adjustments do not have to be made because only one instance of Cassandra is running and the IP addresses are assigned via Docker.

The backend is accessible under `http://localhost:8080/`.
Apache Cassandra is accessible under port `9042` and the credentials are both password and username `cassandra`.
