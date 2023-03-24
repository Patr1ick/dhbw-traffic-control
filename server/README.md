# Server

## Run the server

### Prerequisites

You need to have the following installed:

-   Go

Additionally, you need a running Apache Cassandra instance (listening on Port 9042)

### Build and run

```bash
# Build the server
go build main.go
# Start the server
main.exe -x 10 -y 10 -z 2
```

The server will start and listen to port 8080. It will connect to Apache Cassandra on the keyspace **traffic_control**. It expects a table named **traffic_area**. The command line arguments are the size of the area.

## Docker

```bash
# Build docker container
docker build -t traffic_control .
# Start the server
docker run -p 8080:8080 traffic_control
```
