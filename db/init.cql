CREATE keyspace traffic_control WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

CREATE TABLE traffic_area
(
    x     int,
    y     int,
    z     int,
    value int,
    PRIMARY KEY ((y, z, x))
);
// TODO