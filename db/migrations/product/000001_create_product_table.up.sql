CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS product(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR (50) NOT NULL,
    description VARCHAR (500) NOT NULL,
    price integer NOT NULL,
    counter bigint NOT NULL
    );

CREATE TABLE IF NOT EXISTS product_images (
    id serial primary key,
    product_id uuid not null references product(id),
    url text not null
);