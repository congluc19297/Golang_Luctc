-- CREATE ENUM type
CREATE TYPE type AS ENUM ('vip', 'normal');
CREATE TYPE role AS ENUM ('admin', 'user');

-- CREATE TABLE customers
CREATE TABLE IF NOT EXISTS customers (
    id          SERIAL PRIMARY KEY,
    type_user   type DEFAULT 'normal',
    role_user   role DEFAULT 'user',
    fullname    TEXT NOT NULL,
    username    TEXT NOT NULL,
    password    TEXT NOT NULL,
    email       TEXT NOT NULL,
    phone       VARCHAR(20) NOT NULL
);

-- CREATE TABLE payment_gateway
CREATE TABLE IF NOT EXISTS payment_gateway (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- CREATE TABLE gateway
CREATE TABLE IF NOT EXISTS gateway (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- CREATE TABLE rules
CREATE TABLE IF NOT EXISTS rules (
   id SERIAL PRIMARY KEY,
   gateway_id INT references gateway(id),
   regex TEXT NOT NULL,
   disable BOOLEAN DEFAULT true NOT NULL
);

-- CREATE TABLE tx
CREATE TYPE state AS ENUM ('new', 'pending', 'timeout', 'confirmed')
CREATE TABLE IF NOT EXISTS tx (
    id SERIAL PRIMARY KEY,
    user_id     INT references customers(id),
    method_id   INT references payment_gateway(id),
    gateway_id  INT references gateway(id),
    tx_state    state DEFAULT 'new',
    real_money  money NOT NULL,
    expect_money money NOT NULL,
    account     VARCHAR(26),
    
);

