CREATE TABLE customers (
    customer_id varchar(40) NOT NULL PRIMARY KEY,
    name text NOT NULL,
    phone varchar(20) NOT NULL,
    zip varchar(10) NOT NULL,
    city text NOT NULL,
    address text NOT NULL,
    region text NOT NULL,
    email text NOT NULL
);

CREATE TABLE items (
    chrt_id varchar(40) NOT NULL PRIMARY KEY,
    price numeric(6, 2) NOT NULL,
    rid varchar(30) NOT NULL,
    name text NOT NULL,
    sale int CHECK ( sale >= 0 AND sale <= 100 ),
    size varchar(10),
    nm_id int,
    brand text
);

CREATE TABLE orders (
    order_uid varchar(40) NOT NULL PRIMARY KEY,
    customer_id varchar(40) NOT NULL,
    track_number varchar(14) NOT NULL,
    entry text NOT NULL,
    locale varchar(3) NOT NULL,
    internal_signature text NOT NULL,
    delivery_service text NOT NULL,
    shard_key varchar(10) NOT NULL,
    sm_id int,
    date_created date NOT NULL DEFAULT NOW(),
    oof_shard varchar(10),
    FOREIGN KEY (customer_id)
    REFERENCES customers(customer_id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);

CREATE TABLE order_items (
    order_id varchar(40) NOT NULL,
    item_id varchar(40) NOT NULL,
    status INT NOT NULL,
    PRIMARY KEY (order_id, item_id),
    FOREIGN KEY (order_id)
    REFERENCES orders(order_uid)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
    FOREIGN KEY (item_id)
    REFERENCES items(chrt_id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);

CREATE TABLE payments (
    transaction varchar(40) NOT NULL PRIMARY KEY,
    request_id text,
    currency varchar(5) NOT NULL,
    provider varchar(10) NOT NULL,
    payment_dt int NOT NULL,
    bank text NOT NULL,
    delivery_cost numeric(6, 2) NOT NULL,
    goods_total int NOT NULL,
    custom_fee numeric(6, 2) NOT NULL,
    FOREIGN KEY (transaction) REFERENCES customers(customer_id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);