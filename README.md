# Traffic Control

Project for the subject Distributed System at the Cooperative State University Stuttgart.

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

## Production

### Server

You have to change in the `server/docker-compose.yaml` the following environment variables:

-   `CASSANDRA_SEEDS`: The ip addresses of the other nodes
-   `CASSANDRA_BROADCAST_ADDRESS`: The own ip address

After that you can start the container:

```bash
docker compose up -d
```

Currently you have to initialise the database yourself. Therefore, the script that you execute under `db/init.cql` must be executed on the database. The server will connect to it automatically.

### Client

tbd.
