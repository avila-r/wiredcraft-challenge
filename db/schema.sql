create extension if not exists "pgcrypto";

create table users if not exists (
    id uuid primary key default gen_random_uuid(),
    name varchar(50) not null,
    dob date not null,
    description text not null,
    created_at timestamp not null default current_timestamp
);

create table user_address if not exists (
    id uuid primary key default gen_random_uuid(),
    user_id uuid not null,

    address_line1 varchar(100) not null,
    address_line2 varchar(100),
    city varchar(50) not null,
    state varchar(50) not null,
    postal_code varchar(20) not null,
    country varchar(50) not null,
    created_at timestamp not null default current_timestamp,

    foreign key (user_id) references users(id)
);
