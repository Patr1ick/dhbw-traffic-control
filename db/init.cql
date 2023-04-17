CREATE keyspace traffic_control WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

CREATE TABLE traffic_control.clients
(
    x  int,
    z  int,
    y  int,
    id UUID,
    PRIMARY KEY ((x, z, y))
);

