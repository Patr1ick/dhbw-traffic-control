# Server

## Production

To use the server for production the [docker-compose.yaml](https://github.com/Patr1ick/dhbw-traffic-control/blob/main/server/docker-compose.yaml) can be used.

### Prerequisites

Note that the IP addresses for YugabyteDB need to be changed for the database to work! You have to change the following things **for each node individually**:

| File and Line                                                                                                                  | Description                                                                                                                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [docker-compose.yaml (Line 7)](https://github.com/Patr1ick/dhbw-traffic-control/blob/main/server/docker-compose.yaml#L7)       | Defines the IP-address of the db for the server. Set this value to the current nodes IP-address.                                                                                                                    |
| [docker-compose.yaml (Line 20)](https://github.com/Patr1ick/dhbw-traffic-control/blob/main/server/docker-compose.yaml#L20)     | `master_addresses` defines all YugabyteDB master IP-addresses with the port `7100` in a comma seperated list. Replace with your IP-Address of all nodes where the master is running (for three nodes three master). |
| [docker-compose.yaml (Line 21)](https://github.com/Patr1ick/dhbw-traffic-control/blob/main/server/docker-compose.yaml#L21)     | `rpc_bind_adresses` defines the IP-address the master listens to. Replace here the IP-address with the current node address (with the port `7100`).                                                                 |
| [docker-compose.yaml (Line 24, 25)](https://github.com/Patr1ick/dhbw-traffic-control/blob/main/server/docker-compose.yaml#L24) | Defines the region (datacenter) and zone of the master. Default is here `raspi1` for region and `node1` for zone. Increment the number for each node.                                                               |
| [docker-compose.yaml (Line 41)](https://github.com/Patr1ick/dhbw-traffic-control/blob/main/server/docker-compose.yaml#L41)     | `tserver_master_addrs` defines all YugabyteDB master IP-addresses with the port `7100` in a comma seperated list. Replace with your IP-Address of all nodes where the master is running.                            |
| [docker-compose.yaml (Line 42)](https://github.com/Patr1ick/dhbw-traffic-control/blob/main/server/docker-compose.yaml#L42)     | `rpc_bind_adresses` defines the IP-address the tserver listens to. Replace the IP-address with the current nodes address (with the port `9100`).                                                                    |
| [docker-compose.yaml (Line 43)](https://github.com/Patr1ick/dhbw-traffic-control/blob/main/server/docker-compose.yaml#L43)     | `cql_proxy_bind_adress` defines the address and port where a cql-client can access the db. Replace the IP-address with the current nodes address (with the port `9042`).                                            |
| [docker-compose.yaml (Line 45, 46)](https://github.com/Patr1ick/dhbw-traffic-control/blob/main/server/docker-compose.yaml#L46) | Defines the region (datacenter) and zone of the tserver. Default is here `raspi1` for region and `node1` for zone. Increment the number for each node.                                                              |

### Start the docker container

After that, the docker compose can be started with the following command.

```bash
# Build if needed
sudo docker compose build
# Start container
sudo docker compose up -d
```

The following ports are used and need to be available:
| Port | Description |
|--------|---------------------------------------|
| `7000` | Webserver of YugabyteDB master |
| `7100` | Internode RPC communication (master) |
| `8080` | REST-API of the backend |
| `9000` | Webserver of the YugabyteDB tserver |
| `9100` | Internode RPC communication (tserver) |
| `9042` | Client API for CQL |

The YugabyteDB masters will select a leader according to the Raft concensus. At least two masters need to be available for this proccess. When a leader is elected the server will be able to connect to the database. The server will initialised the keyspace and table if they are not existing. After that the server can used under the port as defined earlier.

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
