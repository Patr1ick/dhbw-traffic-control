# Traffic Control

Project for the subject Distributed Systems at the Cooperative State University Stuttgart.

## Production

Production means the operation of the system for a distributed system. The system is intended for three nodes (servers) and one client, as required in the requirements.

### Requirements

#### Client

-   Docker (for Windows via Docker Desktop and WSL)
-   [Go](https://go.dev/dl/)

#### Server

-   Docker

### Execution

You can find the detailed description to run

-   [server](https://github.com/Patr1ick/dhbw-traffic-control/blob/main/server/README.md)
-   [client](https://github.com/Patr1ick/dhbw-traffic-control/blob/main/client/README.md)

## Local Development Setup

The local development setup should be **only** used for local testing.
To set up the local development environment, simply set up the corresponding docker compose.

```bash
# Build docker container
docker compose -f docker-compose-dev.yaml build
# Start docker container
docker compose -f docker-compose-dev.yaml up -d
```

This starts the Docker containers for the server and YugabyteDB. In the local setup, Caddy is currently not considered, but can be integrated if the IP-addresses in the `Caddyfile` are changed to the local setup. Adjustments do not have to be made because only one YB-Master and YB-TServer is running and the IP-addresses are assigned via Docker.

The backend is accessible under `http://localhost:8080/`.
The CQL API of YugabyteDB is accessible under the port `9042`.

## License

Licensed under [MIT-License](https://github.com/Patr1ick/dhbw-traffic-control/blob/main/LICENSE).
