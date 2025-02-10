create table accounts(
    id bigint PRIMARY KEY,
    name varchar  NOT NULL,
    email varchar NOT NULL,
    password varchar NOT NULL
);
create type category_type AS enum('expense','income');
create table category (
    id bigint PRIMARY KEY, 
    name varchar NOT NULL,
    category category_type  NOT NULL,
    accounts_id bigint REFERENCES accounts(id) ON DELETE CASCADE
);
create table wallet (
    id bigint PRIMARY KEY NOT NULL,
    name varchar NOT NULL,
    balance float NOT NULL,
    accounts_id bigint REFERENCES accounts(id) ON DELETE CASCADE 
);

create type interval_type AS enum('daily','weekly','monthly');

create table budget_plan (
    id bigint PRIMARY KEY,
    name varchar,
    amount float,
    category_id bigint REFERENCES category(id) ON DELETE CASCADE, 
    accounts_id bigint REFERENCES accounts(id) ON DELETE CASCADE NOT NULL, 
    interval interval_type
);
create table cashflow (
    id bigint PRIMARY KEY,
    amount float,
    name varchar ,
    wallet_id bigint REFERENCES wallet(id) ON DELETE CASCADE NOT NULL,
    budget_plan_id  bigint REFERENCES budget_plan(id) ON DELETE CASCADE NOT NULL,
    category_id bigint REFERENCES category(id) ON DELETE CASCADE NOT NULL
);

