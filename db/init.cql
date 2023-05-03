CREATE keyspace traffic_control;

CREATE TABLE traffic_control.clients ( x int, y int, z int, id uuid PRIMARY KEY ) WITH transactions = {'enabled': 'true'};
CREATE UNIQUE INDEX traffic_control_coordinates ON traffic_control.clients(x, y, z);