# Client

## With Caddy (Production)

### Prerequisites

Caddy is used as load balancer and reverse proxy to distribute the requests to the backend. The default port that Caddy listens to is `9000`. The port or host may have to be changed in the `client.go` at the following positions:

-   [client.go (Line 79)](https://github.com/Patr1ick/dhbw-traffic-control/blob/main/client/logic/client.go#L79)
-   [client.go (Line 132)](https://github.com/Patr1ick/dhbw-traffic-control/blob/main/client/logic/client.go#L132)

Additionally, the IP-addresses of the backend have to be changed:

-   [Caddyfile (Line, 4,5,6)](https://github.com/Patr1ick/dhbw-traffic-control/blob/main/client/caddy/Caddyfile#L4)

### Start Caddy

```bash
# Build docker image
sudo docker compose build
# Start Caddy
sudo docker compose up -d
```

### Run client

In order to run the client you do not need to build it (but it is possible).

```bash
go run main.go
```

## Without Caddy

For the local development setup you do not need the Caddy as a load balancer. In order to access the backend you may have to change the host and/or the port (The backend port is per default `8080`). The client can be started with the following command.

```bash
go run main.go
```
