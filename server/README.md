# Server

The server manages the requests from the client. It consists of the backend and the database YugabyteDB.

## Production

To use the server for production the [docker-compose.yaml](https://github.com/Patr1ick/dhbw-traffic-control/blob/main/server/docker-compose.yaml) can be used.

### Prerequisites

In the `docker-compose.yaml` the IP-addresses for the nodes are `192.168.178.47`, `192.168.178.48`, `192.168.178.49`. If you want to use them aswell you still need to adjust the `docker-compose.yaml` **for each node individually**.
If you want to use other IP-adresses you need to change every field which is listed below.

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

After you changed the `docker-compose.yaml` for one node, the docker container can be started with the following command.

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

#### Clean-Up Database

You need install ycqlsh according to the installation tutorial [here](https://docs.yugabyte.com/preview/admin/ycqlsh/).

```bash
# Access database
ycqlsh -u cassandra -p cassandra <ip>:<port>
# Remove all rows
TRUNCATE traffic_control.clients;
```

## Local Development Setup

### Requirements

A running instance of YugabyteDB has to be accessible under the port `9042` (default ycql port). You can use the [docker-compose-dev.yaml](https://github.com/Patr1ick/dhbw-traffic-control/blob/main/docker-compose-dev.yaml) to setup the YugabyteDB (and/or the server).

### Build and run

```bash
# Build
go build main.go
# Start the server (for windows: .\main.exe ...)
main.exe <arguments>
```

| Argument          | Description                                                     | required |
| ----------------- | --------------------------------------------------------------- | -------- |
| `-x`, `--width`   | Width of the field                                              | `true`   |
| `-y`, `--height`  | Height of the field                                             | `true`   |
| `-z`, `--depth`   | Depth of the field (how many clients per positions are allowed) | `true`   |
| `-a`, `--address` | The address of the Yugabyte Database                            | `true`   |

The server will start and listen to port **8080**.

### Docker

Alternatively, the docker image can be build and started as followed. Note that the settings in the docker container is in Release mode which means the webserver is not debugging.

```bash
# Build docker container
docker build -t traffic_control .
# Start the server (insert your address)
docker run -p 8080:8080 -e DATABASE_ADDRESS=<your-database-here> traffic_control
```
