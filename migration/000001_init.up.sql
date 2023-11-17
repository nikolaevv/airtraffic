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

create table if not exists seats (
    id serial primary key,
    seat_no varchar(5) not null,
    aircraft_id int not null
);

create table if not exists aircrafts (
    id serial primary key,
    name varchar(255) not null
);

create table if not exists boarding_passes (
    id serial primary key,
    flight_id int not null,
    seat_id int not null
);

create table if not exists bookings (
    id serial primary key,
    created_at timestamp default now(),
    total_amount float not null
);

create table if not exists flights (
    id serial primary key,
    scheduled_departure timestamp not null,
    scheduled_arrival timestamp not null,
    departure_airport_id int not null,
    arrival_airport_id int not null,
    status flight_status not null,
    aircraft_code varchar(4),
    actual_departure timestamp,
    actual_arrival timestamp
);

create table if not exists tickets (
    id serial primary key,
    booking_id int not null,
    passenger_id int not null
);

create table if not exists ticket_flights (
    id serial primary key,
    flight_id int not null,
    amount float not null
);
