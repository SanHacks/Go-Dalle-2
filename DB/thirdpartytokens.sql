create table thirdpartytokens
(
    id         int auto_increment
        primary key,
    service_id int                                 null,
    user_id    int                                 null,
    token      varchar(255)                        null,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
);

