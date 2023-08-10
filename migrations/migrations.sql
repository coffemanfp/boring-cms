CREATE TABLE IF NOT EXISTS client (
    id serial not null unique,
    name varchar,
    surname varchar,
    username varchar not null unique,
    created_at timestamp,
    password varchar,

    primary key (id)
);

CREATE TABLE IF NOT EXISTS product (
    id serial not null unique,
    client_id integer not null,
    guide_number varchar not null unique,
    type varchar not null,
    joined_at timestamp not null,
    delivered_at timestamp not null,
    shipping_price numeric(19, 5) not null,
    vehicle_plate varchar not null,
    port integer,
    vault integer,

    primary key (id)
);