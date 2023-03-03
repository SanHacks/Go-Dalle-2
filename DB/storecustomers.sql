create table storecustomers
(
    id          int auto_increment
        primary key,
    name        varchar(255)                        null,
    email       varchar(255)                        null,
    password    varchar(255)                        null,
    national_id varchar(255)                        null,
    phone       varchar(255)                        null,
    address     varchar(255)                        null,
    created_at  timestamp default CURRENT_TIMESTAMP not null,
    updated_at  timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
);

