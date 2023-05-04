# Client

## With Caddy (Production)

### Prerequisites

Caddy is used as load balancer and reverse proxy to distribute the requests to the backend. The default port that Caddy listens to is `9000`. Additionally, the IP-addresses of the backend may have to be changed:

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
To start the script you can use the following command.

```bash
go run main.go <arguments>
```

| Argument | Description                                                                                        | required |
| -------- | -------------------------------------------------------------------------------------------------- | -------- |
| `-x`     | The number of vehicles to be generated. The number must not be less than 1 and not more than 1000. | `true`   |
| `-c`     | The address to Caddy. Should be: `localhost:<port>` or `<ip>:<port>`                               | `true`   |

## Without Caddy (Local Development)

For the local development setup you do not need the Caddy as a load balancer. In order to access the backend you may have to change the host and/or the port (The backend port is per default `8080`). The client can be started with the following command. The arguments are the same as above.

```bash
go run main.go <arguments>
```
