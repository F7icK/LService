CREATE TABLE users (
    id bigserial not null primary key,
    name varchar not null,
    surname varchar not null,
    telephone varchar not null unique
)