create table movies(
    movie_id serial primary key,
    title varchar(150) not null,
    description varchar(1000) not null,
    release_date date not null,
    rating numeric(2, 1) not null
);

create table actors(
    actor_id serial primary key,
    name varchar(100) not null,
    gender varchar(10) check(gender in ('male', 'female')) not null,
    birth_date date not null
);

create table actors_movie(
    actor_id integer references actors(actor_id),
    movie_id integer references movies(movie_id),
    constraint unique_actor_movie unique (actor_id, movie_id)
);

create table user_roles(
    role_id serial primary key,
    role_title varchar(20) not null unique
);

create table users(
    user_id serial primary key,
    login varchar(150) not null,
    password varchar(255) not null,
    role varchar(20) not null references user_roles(role_title)
);

INSERT INTO
    actors (name, gender, birth_date)
VALUES
    ('Брэд Питт', 'male', '1963-12-18'),
    ('Анджелина Джоли', 'female', '1975-06-04'),
    ('Леонардо ДиКаприо', 'male', '1974-11-11'),
    ('Дженнифер Лоуренс', 'female', '1990-08-15'),
    ('Том Хэнкс', 'male', '1956-07-09');

INSERT INTO
    movies (title, description, release_date, rating)
VALUES
    (
        'Бойцовский клуб',
        'Офисный работник, страдающий бессонницей, и небрежный изготовитель мыла создают подпольный бойцовский клуб, который превращается в нечто большее.',
        '1999-10-15',
        8.8
    ),
    (
        'Мистер и миссис Смит',
        'Скучающая супружеская пара с удивлением узнает, что они оба наемные убийцы, нанятые конкурирующими агентствами для убийства друг друга.',
        '2005-06-10',
        6.5
    ),
    (
        'Титаник',
        'Семнадцатилетний аристократ влюбляется в доброго, но бедного художника на борту роскошного, обреченного на гибель R.M.S. Титаник.',
        '1997-12-19',
        7.8
    ),
    (
        'Голодные игры',
        'Кэтнисс Эвердин добровольно заменяет свою младшую сестру в Голодных играх: телевизионном соревновании, в котором по два подростка из каждого из двенадцати округов Панема выбираются случайным образом, чтобы сразиться насмерть.',
        '2012-03-23',
        7.2
    ),
    (
        'Форрест Гамп',
        'Президентства Кеннеди и Джонсона, события Вьетнама, Уотергейт и другие исторические события разворачиваются с точки зрения алабамского человека с IQ 75, единственное желание которого - вновь встретиться со своей возлюбленной из детства.',
        '1994-07-06',
        8.8
    );

INSERT INTO
    actors_movie (actor_id, movie_id)
VALUES
    (1, 1),
    (2, 2),
    (1, 2),
    (3, 3),
    (4, 4),
    (5, 5);