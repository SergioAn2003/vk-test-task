create table users(
    id serial primary key,
    name varchar(255) not null,
    age integer not null,
    is_married boolean
);