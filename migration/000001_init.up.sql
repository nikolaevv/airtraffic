CREATE TYPE if not exists flight_status AS ENUM ('SCHEDULED', 'ONTIME', 'DELAYED', 'CANCELLED', 'DEPARTED', 'ARRIVED');

create table if not exists flights (
    id serial primary key,
    scheduled_departure timestamp,
    scheduled_arrival timestamp,
    status flight_status,
    aircraft_code varchar(4),
    actual_departure timestamp,
    actual_arrival timestamp
);
