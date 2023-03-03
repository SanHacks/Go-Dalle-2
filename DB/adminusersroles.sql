create table adminusersroles
(
    id         int auto_increment
        primary key,
    user_id    int                                 null,
    role_id    int                                 null,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
);

