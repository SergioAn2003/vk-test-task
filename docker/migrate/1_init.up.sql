create table actors(
    id serial primary key,
    name varchar(100) not null,
    gender varchar(10) check(gender in ('male', 'female')) not null,
    birth_date date not null
);

INSERT INTO
    actors (name, gender, birth_date)
VALUES
    ('Brad Pitt', 'male', '1963-12-18'),
    ('Angelina Jolie', 'female', '1975-06-04'),
    ('Leonardo DiCaprio', 'male', '1974-11-11'),
    ('Jennifer Lawrence', 'female', '1990-08-15'),
    ('Tom Hanks', 'male', '1956-07-09');