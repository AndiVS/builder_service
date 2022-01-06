CREATE TABLE IF NOT EXISTS buildings
(
    id_user    UUID PRIMARY KEY,
    id_planet  UUID,
    type       VARCHAR(64),
    name       varchar(64),
    update     boolean,
    start_time timestamp,
    time_need  timestamp
);

CREATE TABLE IF NOT EXISTS shipyard
(
    id_user    UUID PRIMARY KEY,
    id_planet  UUID,
    type       VARCHAR(64),
    name       varchar(64),
    start_time timestamp,
    time_need  timestamp

);

CREATE TABLE IF NOT EXISTS laboratory
(
    id_user    UUID PRIMARY KEY,
    id_planet  VARCHAR(64),
    type       VARCHAR(64),
    name       varchar(64),
    start_time timestamp,
    time_need  timestamp
);
