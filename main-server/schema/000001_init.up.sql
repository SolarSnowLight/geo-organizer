CREATE TABLE users
(
    id            serial       primary key not null unique,
    uuid          varchar(36)              not null unique,
    email         varchar(255)             not null unique,
    password      varchar(255)             not null
);

CREATE TABLE users_data
(
    id                  serial primary key not null unique,
    name                varchar(255)       not null,
    surname             varchar(255)       not null,
    date_registration   date               not null,
    users_id int references users (id) on delete cascade not null
);

CREATE TABLE tokens
(
    id                  serial primary key not null unique,
    access_token        varchar(512)       not null,
    refresh_token       varchar(512)       not null,
    users_id int references users (id) on delete cascade not null
);

CREATE TABLE roles
(
    id              serial primary key not null unique,
    uuid            varchar(36)        not null unique,
    value           varchar(255)       not null unique,
    description     varchar(255)       not null,
    users_id int references users (id) on delete cascade
);

CREATE TABLE roles_modules
(
    id      serial primary key  not null unique,
    users   boolean             not null,
    admins  boolean             not null,
    devs    boolean             not null,
    roles_id int references roles (id) on delete cascade not null
);

CREATE TABLE users_roles
(
    id              serial primary key not null unique,
    users_id int references users (id) on delete cascade not null,
    roles_id int references roles (id) on delete cascade not null
);

CREATE TABLE activations
(
    id                  serial primary key not null unique,
    is_activated        boolean            not null,
    activation_link     varchar(255)       not null,
    users_id int references users (id) on delete cascade not null
);


CREATE TABLE todo_lists
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255)
);

CREATE TABLE users_lists
(
    id      serial                                           not null unique,
    user_id int references users (id) on delete cascade      not null,
    list_id int references todo_lists (id) on delete cascade not null
);

CREATE TABLE todo_items
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255),
    done        boolean      not null default false
);


CREATE TABLE lists_items
(
    id      serial                                           not null unique,
    item_id int references todo_items (id) on delete cascade not null,
    list_id int references todo_lists (id) on delete cascade not null
);