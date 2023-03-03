create table storeorders
(
    id          int auto_increment
        primary key,
    customer_id int                                 null,
    total       decimal(10, 2)                      null,
    created_at  timestamp default CURRENT_TIMESTAMP not null,
    updated_at  timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
);

