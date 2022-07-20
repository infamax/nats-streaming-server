CREATE TABLE orders (
    id serial PRIMARY KEY,
    order_uid varchar(40) NOT NULL,
    track_number varchar(40) NOT NULL,
    entry text,
    locale text,
    internal_signature text,
    customer_id text NOT NULL ,
    delivery_service text,
    shardkey text,
    sm_id integer NOT NULL,
    date_created date DEFAULT NOW(),
    oof_shard text
);

CREATE TABLE items (
    id serial PRIMARY KEY,
    chrt_id integer NOT NULL,
    track_number varchar(40) NOT NULL,
    price numeric(6, 2),
    rid text,
    name text,
    sale integer,
    size integer,
    total_price numeric(6, 2),
    nm_id integer NOT NULL,
    brand text,
    status integer,
    FOREIGN KEY (track_number)
    REFERENCES orders(track_number)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);

CREATE TABLE delivery (
    id serial PRIMARY KEY,
    order_id integer,
    name text,
    phone text,
    zip text,
    city text,
    address text,
    region text,
    email text,
    FOREIGN KEY (order_id)
    REFERENCES orders(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);

CREATE TABLE payments (
    id serial PRIMARY KEY,
    transaction varchar(40) NOT NULL,
    request_id text NOT NULL,
    currency varchar(5),
    provider text,
    amount integer,
    payment_id integer NOT NULL,
    bank text,
    delivery_coast numeric(6, 2),
    goods_total integer,
    custom_fee integer,
    FOREIGN KEY (transaction)
    REFERENCES orders(order_uid)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);

CREATE TABLE invalid_messages(
    id serial PRIMARY KEY,
    data jsonb
);

