# Server

## Production

To use the server for production the [docker-compose.yaml](https://github.com/Patr1ick/dhbw-traffic-control/blob/main/server/docker-compose.yaml) can be used.

### Prerequisites

Note that the IP addresses for YugabyteDB need to be changed for the database to work! You have to change the following things **for each node individually**:

-   [docker-compose.yaml (Line 22)](https://github.com/Patr1ick/dhbw-traffic-control/blob/main/server/docker-compose.yaml#L22): The `master_addresses` specifies all masters of Yugabyte which runs on all (three) nodes.
-   [docker-compose.yaml (Line 23)](https://github.com/Patr1ick/dhbw-traffic-control/blob/main/server/docker-compose.yaml#L23): The `rpc_bind_addresses` specifies the IP-address the Yugabyte master instance is listening to.
-   [docker-compose.yaml (Line 42)](https://github.com/Patr1ick/dhbw-traffic-control/blob/main/server/docker-compose.yaml#L22): The `tserver_master_addrs` specifies in a comma-seperated list all IP-addresses for the `yb-master` which runs on all (three) nodes.
-   [docker-compose.yaml (Line 23)](https://github.com/Patr1ick/dhbw-traffic-control/blob/main/server/docker-compose.yaml#L23): The `rpc_bind_addresses` specifies the IP-address the Yugabyte tserver instance is listening to.

### Start the docker container

After that, the docker compose can be started with the following command.

```bash
sudo docker compose up -d
```

This starts the server and the database container.

[//]: # "Setup database further: Init keyspace and table"

## Local Development Setup

### Requirements

A running instance of YugabyteDB has to be accessible under the port `9042` (default ycql port). Additionally, the keyspace `traffic_control` and the table `clients` has to be existing (To setup the database the script [init.cql](https://github.com/Patr1ick/dhbw-traffic-control/blob/main/db/init.cql) can be used.)

### Build and run

```bash
# Build
go build main.go
# Start the server (for windows: .\main.exe ...)
main.exe <arguments>
```

| Argument          | Description                                                     | required |
| ----------------- | --------------------------------------------------------------- | -------- |
| `-x`              | Width of the field                                              | `true`   |
| `-y`              | Height of the field                                             | `true`   |
| `-z`              | Depth of the field (how many clients per positions are allowed) | `true`   |
| `-a`, `--address` | The address of the Yugabyte Database                            | `true`   |

The server will start and listen to port **8080**.

### Docker

```bash
# Build docker container
docker build -t traffic_control .
# Start the server (insert your address)
docker run -p 8080:8080 -e DATABASE_ADDRESS=<your-database-here> traffic_control
```
