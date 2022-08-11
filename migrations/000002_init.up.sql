CREATE TABLE IF NOT EXISTS orders
(
order_uid          TEXT PRIMARY KEY,
track_number       TEXT NOT NULL,
entry              TEXT NOT NULL,
locale             TEXT NOT NULL,
internal_signature TEXT NOT NULL,
customer_id        TEXT NOT NULL,
delivery_service   TEXT NOT NULL,
shardkey           TEXT NOT NULL,
sm_id              INT  NOT NULL,
date_created       TIMESTAMPTZ NOT NULL,
oof_shard          TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS delivery
(
id        SERIAL PRIMARY KEY,
name      TEXT   NOT NULL
phone     TEXT   NOT NULL
zip       TEXT   NOT NULL
city      TEXT   NOT NULL
address   TEXT   NOT NULL
region    TEXT   NOT NULL
email     TEXT   NOT NULL
order_uid TEXT   NOT NULL

FOREIGN KEY (order_uid) REFERENCES orders(order_uid) ON DELETE CASCADE 
);

CREATE TABLE IF NOT EXISTS payment
(
    
)