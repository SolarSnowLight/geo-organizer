/* Таблицы верификации пользователя и управления доступом на основе ролей (RBAC)*/
/* ---------------------------------------------------------------------------- */

/* Таблица идентификации пользователя внутри системы (main) */
CREATE TABLE users
(
    id            serial       primary key not null,
    uuid          varchar(36)              not null unique,
    email         varchar(255)             not null unique,
    "password"    varchar(1024)             not null
);

/* Таблица токена доступа и обновления для верификации пользователя */
CREATE TABLE tokens
(
    id                  serial primary key not null,
    access_token        varchar(1024)       not null,
    refresh_token       varchar(1024)       not null,
    users_id            int references users (id) on delete cascade not null
);

/* Таблица доменов, характеризующие различные подсистемы характеризующие отдельные контейнеры */
CREATE TABLE domains
(
    id                  serial primary key  not null,
    uuid                varchar(36)         not null unique,
    "value"             varchar(10)         not null,
    "description"       varchar(512),
    users_id            int references users (id) on delete cascade not null
);

/* Таблица всех существующих ролей */
CREATE TABLE roles
(
    id              serial primary key not null,
    uuid            varchar(36)        not null unique,
    "value"         varchar(255)       not null unique,
    "description"   varchar(512),
    users_id        int references users (id) on delete cascade,
    domains_id      int references domains (id) on delete cascade
);

/* Таблица всех супер администраторов */
CREATE TABLE super_admins
(
    id              serial primary key not null,
    users_id        int references users (id) on delete cascade
);

/* Таблица атрибутов */
CREATE TABLE attributes
(
    id              serial primary key not null,
    "value"         varchar(255)       not null unique,
    "description"   varchar(512),
    domains_id      int references domains (id) on delete cascade
);

/* Таблица связывающая роли и атрибуты */
CREATE TABLE roles_attributes
(
    id              serial primary key not null,
    roles_id        int references roles (id) on delete cascade,
    attributes_id   int references attributes (id) on delete cascade
);

/* Таблица всех существующих модулей */
CREATE TABLE modules
(
    id              serial primary key not null,
    "value"         varchar(255)       not null unique,
    "description"   varchar(512),
    domains_id      int references domains (id) on delete cascade
);

/* Таблица связывающая роли с модулями */
CREATE TABLE roles_modules
(
    id              serial primary key not null unique,
    modules_id      int references modules (id) on delete cascade not null,
    roles_id        int references roles (id) on delete cascade not null
);

/* Таблица связывающая пользователей с конкретными ролями */
CREATE TABLE users_roles
(
    id              serial primary key not null unique,
    users_id        int references users (id) on delete cascade not null,
    roles_id        int references roles (id) on delete cascade not null
);

/* Таблица активации аккаунта каждого пользователя */
CREATE TABLE activations
(
    id                  serial primary key not null unique,
    is_activated        boolean            not null,
    activation_link     varchar(255)       not null,
    users_id            int references users (id) on delete cascade not null
);

/* Пользовательские данные */
CREATE TABLE users_data
(
    id                  serial primary key not null unique,
    "data"              jsonb              not null,
    date_registration   date               not null,

    users_id            int references users (id) on delete cascade not null
);

/* Таблица существующих типов аутентификации пользователя */
CREATE TABLE auth_types
(
    id                  serial primary key not null unique,
    uuid                varchar(36)        not null unique,
    "value"             varchar(100)       not null
);

/* Таблица связывания существующих типов аутентификации с пользователями */
CREATE TABLE users_auth_types
(
    id                  serial primary key not null unique,
    users_id            int references users (id) on delete cascade not null,
    auth_types_id       int references auth_types (id) on delete cascade not null
);