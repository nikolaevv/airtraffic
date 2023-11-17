CREATE TYPE if not exists flight_status AS ENUM ('SCHEDULED', 'ONTIME', 'DELAYED', 'CANCELLED', 'DEPARTED', 'ARRIVED');

create table if not exists passengers (
    id serial primary key,
    fullname varchar(255) not null,
    email varchar(255)
);

create table if not exists airports (
    id serial primary key,
    code varchar(3) not null,
    name varchar(255) not null,
    city varchar(255) not null,
    longitude float not null,
    latitude float not null,
    timezone varchar(255),
);

create table if not exists aircrafts (
    id serial primary key,
    name varchar(255) not null
);

create table if not exists seats (
    id serial primary key,
    seat_no varchar(5) not null,
    aircraft_id int references aircrafts(id)
);

create table if not exists flights (
    id serial primary key,
    scheduled_departure timestamp not null,
    scheduled_arrival timestamp not null,
    departure_airport_id int references airports(id),
    arrival_airport_id int references airports(id),
    status flight_status not null,
    aircraft_id varchar(4) references aircrafts(id),
    actual_departure timestamp,
    actual_arrival timestamp
);

create table if not exists boarding_passes (
    id serial primary key,
    flight_id int references flights(id),
    seat_id int references seats(id)
);

create table if not exists bookings (
    id serial primary key,
    created_at timestamp default now(),
    total_amount float not null
);

create table if not exists tickets (
    id serial primary key,
    booking_id int references bookings(id),
    passenger_id int references passengers(id),
);

create table if not exists ticket_flights (
    id serial primary key,
    flight_id int references flights(id),
    amount float not null
);
