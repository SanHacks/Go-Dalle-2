create table storeorderitems
(
    id         int auto_increment
        primary key,
    order_id   int                                 null,
    product_id int                                 null,
    quantity   int                                 null,
    price      decimal(10, 2)                      null,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
);

