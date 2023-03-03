create table storeorderstatuseshistory
(
    id         int auto_increment
        primary key,
    order_id   int                                 null,
    status_id  int                                 null,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
);

