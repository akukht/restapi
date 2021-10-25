CREATE TABLE users
(
    id serial not null unique,
    login  varchar(50) not null,
    timezone  varchar(50) not null,
    token  varchar(255),
    password  varchar(100) not null
);

CREATE TABLE events
(
    id serial not null unique,
    name  varchar(255) not null,
    description  varchar(255),
    timezone  varchar(50) not null,
    reminding varchar(50)
);

CREATE TABLE users_events
(
    id serial not null unique,
    user_id int references users (id) on delete cascade not null,
    event_id int references events (id) on delete cascade not null
);

CREATE TABLE date_lists(
    id serial not null unique,
    event_id int references events (id) on delete cascade not null,
    year int not null,
    month int not null,
    day int not null,
    hour int not null,
    minutes int not null
);