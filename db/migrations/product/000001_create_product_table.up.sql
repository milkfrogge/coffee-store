CREATE
    EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS category
(
    id   uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(50) NOT NULL UNIQUE
);

INSERT INTO category (id, name)
VALUES ('00000000-0000-0000-0000-000000000001', 'Coffee');
INSERT INTO category (id, name)
VALUES ('00000000-0000-0000-0000-000000000002', 'Drinks');
INSERT INTO category (id, name)
VALUES ('00000000-0000-0000-0000-000000000003', 'Burgers');
INSERT INTO category (id, name)
VALUES ('00000000-0000-0000-0000-000000000004', 'Sweets');

CREATE TABLE IF NOT EXISTS product
(
    id          uuid PRIMARY KEY      DEFAULT uuid_generate_v4(),
    name        VARCHAR(50)  NOT NULL,
    description VARCHAR(500) NOT NULL,
    price       bigint       NOT NULL,
    barista_needed bool NOT NULL default false,
    kitchen_needed bool NOT NULL default false,
    counter     bigint       NOT NULL,
    created_at  timestamp    NOT NULL default now(),
    category    uuid         not null references category (id) on DELETE CASCADE
);

INSERT INTO product (id, name, description, price, category, counter)
VALUES ('00000000-0000-0000-0000-000000000001', 'Latte', 'Description of Latte', 240,
        '00000000-0000-0000-0000-000000000001', 500);

INSERT INTO product (id, name, description, price, category, counter)
VALUES ('00000000-0000-0000-0000-000000000002', 'Cappucino', 'Description of Latte', 250,
        '00000000-0000-0000-0000-000000000001', 500);

INSERT INTO product (id, name, description, price, category, counter)
VALUES ('00000000-0000-0000-0000-000000000003', 'ZXC', 'Description of Latte', 260,
        '00000000-0000-0000-0000-000000000001', 500);

CREATE TABLE IF NOT EXISTS product_images
(
    id         serial primary key,
    product_id uuid not null references product (id) on DELETE CASCADE,
    url        text not null
);

INSERT INTO product_images (product_id, url)
VALUES ('00000000-0000-0000-0000-000000000001',
        'https://upload.wikimedia.org/wikipedia/commons/thumb/9/98/Latte_with_winged_tulip_art.jpg/800px-Latte_with_winged_tulip_art.jpg')