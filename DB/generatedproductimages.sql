create table generatedproductimages
(
    id         int auto_increment
        primary key,
    product_id int                                 null,
    image      varchar(255)                        null,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
);

